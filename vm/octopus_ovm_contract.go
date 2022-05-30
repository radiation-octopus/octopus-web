package vm

import (
	"errors"
	"math/big"
	"octopus/utils"
)

type ContractRef interface {
	Address() utils.Address
}

//合同表示状态数据库中的以太坊合同。它包含合同代码，调用参数。合同执行ContractRef
type Contract struct {
	// CallerAddress是初始化此合同的呼叫者的结果
	CallerAddress utils.Address
	caller        ContractRef
	self          ContractRef

	//jumpdests map[blockchain.Hash]bitvec // Aggregated result of JUMPDEST analysis.
	//analysis  bitvec                 // Locally cached result of JUMPDEST analysis

	Code     []byte
	CodeHash utils.Hash
	CodeAddr *utils.Address
	Input    []byte

	Gas   uint64
	value *big.Int
}

type AccountRef utils.Address

func (ar AccountRef) Address() utils.Address { return (utils.Address)(ar) }

func (c *Contract) UseGas(gas uint64) (ok bool) {
	if c.Gas < gas {
		return false
	}
	c.Gas -= gas
	return true
}

type PrecompiledContract interface {
	RequiredGas(input []byte) uint64  // RequiredPrice calculates the contract gas use
	Run(input []byte) ([]byte, error) // Run runs the precompiled contract
}

type ecrecover struct{}

func (c *ecrecover) RequiredGas(input []byte) uint64 {
	return utils.EcrecoverGas
}
func (c *ecrecover) Run(input []byte) ([]byte, error) {
	const ecRecoverInputLength = 128

	//input = common.RightPadBytes(input, ecRecoverInputLength)
	//// "input" is (hash, v, r, s), each 32 bytes
	//// but for ecrecover we want (r, s, v)
	//
	//r := new(big.Int).SetBytes(input[64:96])
	//s := new(big.Int).SetBytes(input[96:128])
	//v := input[63] - 27
	//
	//// tighter sig s values input homestead only apply to tx sigs
	//if !allZero(input[32:63]) || !crypto.ValidateSignatureValues(v, r, s, false) {
	//	return nil, nil
	//}
	//// We must make sure not to modify the 'input', so placing the 'v' along with
	//// the signature needs to be done on a new allocation
	//sig := make([]byte, 65)
	//copy(sig, input[64:128])
	//sig[64] = v
	//// v needs to be at the end for libsecp256k1
	//pubKey, err := crypto.Ecrecover(input[:32], sig)
	//// make sure the public key is a valid one
	//if err != nil {
	//	return nil, nil
	//}
	//
	//// the first byte of pubkey is bitcoin heritage
	//return common.LeftPadBytes(crypto.Keccak256(pubKey[1:])[12:], 32), nil
	return nil, nil
}

var PrecompiledContractsHomestead = map[utils.Address]PrecompiledContract{
	//blockchain.BytesToAddress([]byte{1}): &ecrecover{},
	//blockchain.BytesToAddress([]byte{2}): &sha256hash{},
	//blockchain.BytesToAddress([]byte{3}): &ripemd160hash{},
	//blockchain.BytesToAddress([]byte{4}): &dataCopy{},
}

func RunPrecompiledContract(p PrecompiledContract, input []byte, suppliedGas uint64) (ret []byte, remainingGas uint64, err error) {
	gasCost := p.RequiredGas(input)
	if suppliedGas < gasCost {
		return nil, 0, errors.New("gas用完")
	}
	suppliedGas -= gasCost
	output, err := p.Run(input)
	return output, suppliedGas, err
}

func (c *Contract) Address() utils.Address {
	return c.self.Address()
}

//返回ovm新合同环境
func NewContract(caller ContractRef, object ContractRef, value *big.Int, gas uint64) *Contract {
	c := &Contract{CallerAddress: caller.Address(), caller: caller, self: object}

	//if parent, ok := caller.(*Contract); ok {
	//	//
	//	c.jumpdests = parent.jumpdests
	//} else {
	//	c.jumpdests = make(map[blockchain.Hash]bitvec)
	//}

	c.Gas = gas

	c.value = value

	return c
}

func (c *Contract) SetCallCode(addr *utils.Address, hash utils.Hash, code []byte) {
	c.Code = code
	c.CodeHash = hash
	c.CodeAddr = addr
}

// GetOp returns the n'th element in the contract's byte array
func (c *Contract) GetOp(n uint64) utils.OpCode {
	if n < uint64(len(c.Code)) {
		return utils.OpCode(c.Code[n])
	}

	return utils.STOP
}
