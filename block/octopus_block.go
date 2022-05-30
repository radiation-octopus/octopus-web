package block

import (
	"encoding/binary"
	"math/big"
	"octopus/utils"
	"sync/atomic"
)

type BlockNonce [8]byte

//区块头结构体
type Header struct {
	ParentHash utils.Hash `autoInjectCfg:"octopus.blockchain.binding.genesis.header.parentHash"` //父hash
	UncleHash  utils.Hash `autoInjectCfg:"octopus.blockchain.binding.genesis.header.uncleHash"`  //叔hash
	//Coinbase    common.Address
	Root        utils.Hash `autoInjectCfg:"octopus.blockchain.binding.genesis.header.root"`        //根hash
	TxHash      utils.Hash `autoInjectCfg:"octopus.blockchain.binding.genesis.header.txhash"`      //交易hash
	ReceiptHash utils.Hash `autoInjectCfg:"octopus.blockchain.binding.genesis.header.receiptHash"` //收据hash
	//Bloom       Bloom
	Difficulty *big.Int `autoInjectCfg:"octopus.blockchain.binding.genesis.header.difficulty"` //难度值
	Number     *big.Int `autoInjectCfg:"octopus.blockchain.binding.genesis.header.number"`     //数量
	GasLimit   uint64   `autoInjectCfg:"octopus.blockchain.binding.genesis.header.gasLimit"`   //gas限制
	GasUsed    uint64   `autoInjectCfg:"octopus.blockchain.binding.genesis.header.gasUsed"`    //gas总和
	Time       uint64   `autoInjectCfg:"octopus.blockchain.binding.genesis.header.time"`       //时间戳
	//Extra       []byte
	MixDigest utils.Hash `autoInjectCfg:"octopus.blockchain.binding.genesis.header.mixDigest"` //mixhash
	Nonce     BlockNonce `autoInjectCfg:"octopus.blockchain.binding.genesis.header.nonce"`     //唯一标识s

	//基本费用
	BaseFee *big.Int `autoInjectCfg:"octopus.blockchain.binding.genesis.header.baseFee"`
}

func (h *Header) Hash() utils.Hash {
	//哈希运算
	//return rlpHash(h)
	return utils.Hash{0}
}

//数据容器
type Body struct {
	Transactions []*Transaction
	Uncles       []*Header
}

type Block struct {
	header       *Header      //区块头信息
	uncles       []*Header    //叔块头信息
	transactions Transactions //交易信息
	// caches
	hash atomic.Value //缓存hash
	size atomic.Value //缓存大小
	td   *big.Int     //交易总难度
}

func (b Block) newGenesis() {

}

//获取交易集
func (b *Block) Transactions() Transactions { return b.transactions }

func (b *Block) Number() *big.Int     { return new(big.Int).Set(b.header.Number) }
func (b *Block) GasLimit() uint64     { return b.header.GasLimit }
func (b *Block) GasUsed() uint64      { return b.header.GasUsed }
func (b *Block) Difficulty() *big.Int { return new(big.Int).Set(b.header.Difficulty) }
func (b *Block) Time() uint64         { return b.header.Time }

func (b *Block) NumberU64() uint64       { return b.header.Number.Uint64() }
func (b *Block) Nonce() uint64           { return binary.BigEndian.Uint64(b.header.Nonce[:]) }
func (b *Block) Root() utils.Hash        { return b.header.Root }
func (b *Block) ParentHash() utils.Hash  { return b.header.ParentHash }
func (b *Block) TxHash() utils.Hash      { return b.header.TxHash }
func (b *Block) ReceiptHash() utils.Hash { return b.header.ReceiptHash }
func (b *Block) UncleHash() utils.Hash   { return b.header.UncleHash }

func (b *Block) BaseFee() *big.Int {
	if b.header.BaseFee == nil {
		return nil
	}
	return new(big.Int).Set(b.header.BaseFee)
}

func (b *Block) Header() *Header { return CopyHeader(b.header) }

// Body returns the non-header content of the block.
func (b *Block) Body() *Body { return &Body{b.transactions, b.uncles} }

func (b *Block) Hash() utils.Hash {
	if hash := b.hash.Load(); hash != nil {
		return hash.(utils.Hash)
	}
	v := b.header.Hash()
	b.hash.Store(v)
	return v
}

func CopyHeader(h *Header) *Header {
	cpy := *h
	if cpy.Difficulty = new(big.Int); h.Difficulty != nil {
		cpy.Difficulty.Set(h.Difficulty)
	}
	if cpy.Number = new(big.Int); h.Number != nil {
		cpy.Number.Set(h.Number)
	}
	if h.BaseFee != nil {
		cpy.BaseFee = new(big.Int).Set(h.BaseFee)
	}
	//if len(h.Extra) > 0 {
	//	cpy.Extra = make([]byte, len(h.Extra))
	//	copy(cpy.Extra, h.Extra)
	//}
	return &cpy
}
