package vm

import (
	"errors"
	"math/big"
	"octopus/block"
	"octopus/consensus"
	"octopus/db"
	"octopus/log"
	"octopus/utils"
)

type BlockContext struct {
	//判断是否有足够的gas转账
	CanTransfer CanTransferFunc
	//转账
	Transfer TransferFunc
	// 获取hash
	GetHash GetHashFunc

	// 区块信息
	Coinbase    utils.Address
	GasLimit    uint64
	BlockNumber *big.Int
	Time        *big.Int
	Difficulty  *big.Int
	BaseFee     *big.Int
	Random      *utils.Hash
}

type (
	// 是否有足够的余额
	CanTransferFunc func(StateDB, utils.Address, *big.Int) bool
	// 交易执行函数
	TransferFunc func(StateDB, utils.Address, utils.Address, *big.Int)
	// 返回第几块的hash
	GetHashFunc func(uint64) utils.Hash
)

// 事务信息
type TxContext struct {
	Origin   utils.Address
	GasPrice *big.Int
}

type OVM struct {
	// 区块配置信息
	Context BlockContext
	TxContext
	// 操作数据库访问配置
	operationdb db.OperationDB
	// 当前调用堆栈深度
	depth int

	//链信息
	//chainConfig *params.ChainConfig
	// 链规则
	//chainRules params.Rules
	// 初始化虚拟机配置选项
	Config Config
	// 整个事务中使用的全局辐射章鱼虚拟机
	interpreter *OVMInterpreter
	// 终止虚拟机调用操作
	abort int32
	// callGasTemp保存当前调用的可用gas
	callGasTemp uint64
}

type ChainContext interface {
	// 共识引擎
	Engine() consensus.Engine

	// 返回其对应hash
	GetHeader(utils.Hash, uint64) *block.Header
}

type StateDB interface {
	CreateAccount(utils.Address)

	SubBalance(utils.Address, *big.Int)
	AddBalance(utils.Address, *big.Int)
	GetBalance(utils.Address) *big.Int

	GetNonce(utils.Address) uint64
	SetNonce(utils.Address, uint64)

	GetCodeHash(utils.Address) utils.Hash
	GetCode(utils.Address) []byte
	SetCode(utils.Address, []byte)
	GetCodeSize(utils.Address) int

	AddRefund(uint64)
	SubRefund(uint64)
	GetRefund() uint64

	GetCommittedState(utils.Address, utils.Hash) utils.Hash
	GetState(utils.Address, utils.Hash) utils.Hash
	SetState(utils.Address, utils.Hash, utils.Hash)

	Suicide(utils.Address) bool
	HasSuicided(utils.Address) bool

	// Exist reports whether the given account exists in state.
	// Notably this should also return true for suicided accounts.
	Exist(utils.Address) bool
	// Empty returns whether the given account is empty. Empty
	// is defined according to EIP161 (balance = nonce = code = 0).
	Empty(utils.Address) bool

	//PrepareAccessList(sender blockchain.Address, dest *blockchain.Address, precompiles []blockchain.Address, txAccesses db.AccessList)
	AddressInAccessList(addr utils.Address) bool
	SlotInAccessList(addr utils.Address, slot utils.Hash) (addressOk bool, slotOk bool)
	// AddAddressToAccessList adds the given address to the access list. This operation is safe to perform
	// even if the feature/fork is not active yet
	AddAddressToAccessList(addr utils.Address)
	// AddSlotToAccessList adds the given (address,slot) to the access list. This operation is safe to perform
	// even if the feature/fork is not active yet
	AddSlotToAccessList(addr utils.Address, slot utils.Hash)

	RevertToSnapshot(int)
	Snapshot() int

	AddLog(*log.OctopusLog)
	AddPreimage(utils.Hash, []byte)

	ForEachStorage(utils.Address, func(utils.Hash, utils.Hash) bool) error
}

func NewOVM(blockCtx BlockContext, txCtx TxContext, operation *db.OperationDB, config Config) *OVM {
	evm := &OVM{
		Context:     blockCtx,
		TxContext:   txCtx,
		operationdb: *operation,
		Config:      config,
	}
	evm.interpreter = NewOVMInterpreter(evm, config)
	return evm
}

func NewOVMBlockContext(header *block.Header, chain ChainContext, author *utils.Address) BlockContext {
	var (
		beneficiary utils.Address
		baseFee     *big.Int
		random      *utils.Hash
	)

	//
	if author == nil {
		beneficiary, _ = chain.Engine().Author(header) // Ignore error, we're past header validation
	} else {
		beneficiary = *author
	}
	if header.BaseFee != nil {
		baseFee = new(big.Int).Set(header.BaseFee)
	}
	if header.Difficulty.Cmp(utils.Big0) == 0 {
		random = &header.MixDigest
	}
	return BlockContext{
		CanTransfer: CanTransfer,
		Transfer:    Transfer,
		GetHash:     GetHashFn(header, chain),
		Coinbase:    beneficiary,
		BlockNumber: new(big.Int).Set(header.Number),
		Time:        new(big.Int).SetUint64(header.Time),
		Difficulty:  new(big.Int).Set(header.Difficulty),
		BaseFee:     baseFee,
		GasLimit:    header.GasLimit,
		Random:      random,
	}
}

func (ovm *OVM) Call(caller ContractRef, addr utils.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
	//深度限制
	if ovm.depth > int(utils.CallCreateDepth) {
		return nil, gas, errors.New("超过最大呼叫深度")
	}
	ovm.Context.Transfer(ovm.operationdb, caller.Address(), addr, value)

	p, isPrecompile := ovm.precompile(addr)
	if isPrecompile {
		ret, gas, err = RunPrecompiledContract(p, input, gas)
	} else {
		// 初始化新合同并设置EVM要使用的代码
		code := ovm.operationdb.GetCode(addr)
		if len(code) == 0 {
			ret, err = nil, nil // gas不变
		} else {
			addrCopy := addr
			// If the account has no code, we can abort here
			// The depth-check is already done, and precompiles handled above
			contract := NewContract(caller, AccountRef(addrCopy), value, gas)
			contract.SetCallCode(&addrCopy, ovm.operationdb.GetCodeHash(addrCopy), code)
			ret, err = ovm.interpreter.Run(contract, input, false)
			gas = contract.Gas
		}
	}
	return nil, 0, err
}

func GetHashFn(ref *block.Header, chain ChainContext) func(n uint64) utils.Hash {
	var cache []utils.Hash

	return func(n uint64) utils.Hash {
		if len(cache) == 0 {
			cache = append(cache, ref.ParentHash)
		}
		if idx := ref.Number.Uint64() - n - 1; idx < uint64(len(cache)) {
			return cache[idx]
		}
		//我们可以从已知的最后一个元素开始迭代
		lastKnownHash := cache[len(cache)-1]
		lastKnownNumber := ref.Number.Uint64() - uint64(len(cache))

		for {
			header := chain.GetHeader(lastKnownHash, lastKnownNumber)
			if header == nil {
				break
			}
			cache = append(cache, header.ParentHash)
			lastKnownHash = header.ParentHash
			lastKnownNumber = header.Number.Uint64() - 1
			if n == lastKnownNumber {
				return lastKnownHash
			}
		}
		return utils.Hash{}
	}
}

func CanTransfer(db StateDB, addr utils.Address, amount *big.Int) bool {
	return db.GetBalance(addr).Cmp(amount) >= 0
}

// Transfer subtracts amount from sender and adds amount to recipient using the given Db
func Transfer(db StateDB, sender, recipient utils.Address, amount *big.Int) {
	db.SubBalance(sender, amount)
	db.AddBalance(recipient, amount)
}

//预编译
func (ovm *OVM) precompile(addr utils.Address) (PrecompiledContract, bool) {
	var precompiles map[utils.Address]PrecompiledContract
	precompiles = PrecompiledContractsHomestead
	p, ok := precompiles[addr]
	return p, ok
}
