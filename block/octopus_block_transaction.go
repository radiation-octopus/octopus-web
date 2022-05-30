package block

import (
	"math/big"
	"octopus/utils"
	"sync/atomic"
	"time"
)

type Transaction struct {
	inner TxData    // 交易共识内容
	time  time.Time // 交易在本地的出现的时间戳

	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

type TxData interface {
	copy() TxData // 复制，初始化所以字段

	chainID() *big.Int
	//accessList() AccessList
	data() []byte
	gas() uint64
	gasPrice() *big.Int
	gasTipCap() *big.Int
	gasFeeCap() *big.Int
	value() *big.Int
	nonce() uint64
	to() *utils.Address

	//rawSignatureValues() (v, r, s *big.Int)
	//setSignatureValues(chainID, v, r, s *big.Int)
}

func (tx *Transaction) Data() []byte { return tx.inner.data() }

// 返回交易的gas限制
func (tx *Transaction) Gas() uint64 { return tx.inner.gas() }

// 返回交易的gas价格
func (tx *Transaction) GasPrice() *big.Int { return new(big.Int).Set(tx.inner.gasPrice()) }

// 返回交易的gas价格上限
func (tx *Transaction) GasTipCap() *big.Int { return new(big.Int).Set(tx.inner.gasTipCap()) }

// 返回交易中每个gas的费用上限
func (tx *Transaction) GasFeeCap() *big.Int { return new(big.Int).Set(tx.inner.gasFeeCap()) }

// 返回交易的金额
func (tx *Transaction) Value() *big.Int { return new(big.Int).Set(tx.inner.value()) }

// 返回交易的发送方账户nonce
func (tx *Transaction) Nonce() uint64 { return tx.inner.nonce() }

// 返回交易的收件人地址
func (tx *Transaction) To() *utils.Address {
	return copyAddressPtr(tx.inner.to())
}

// 返回交易hash
func (tx *Transaction) Hash() utils.Hash {
	if hash := tx.hash.Load(); hash != nil {
		return hash.(utils.Hash)
	}
	var h utils.Hash
	//if tx.Type() == LegacyTxType {
	//	h = rlpHash(tx.inner)
	//} else {
	//	h = prefixedRlpHash(tx.Type(), tx.inner)
	//}
	h = utils.Hash{0}
	tx.hash.Store(h)
	return h
}

type Transactions []*Transaction

//交易消息结构体
type Message struct {
	to        *utils.Address
	from      utils.Address
	nonce     uint64
	amount    *big.Int
	gasLimit  uint64
	gasPrice  *big.Int
	gasFeeCap *big.Int
	gasTipCap *big.Int
	data      []byte
	//accessList AccessList
	isFake bool
}

func NewMessage(from utils.Address, to *utils.Address, nonce uint64, amount *big.Int, gasLimit uint64, gasPrice, gasFeeCap, gasTipCap *big.Int, data []byte, isFake bool) Message {
	return Message{
		from:      from,
		to:        to,
		nonce:     nonce,
		amount:    amount,
		gasLimit:  gasLimit,
		gasPrice:  gasPrice,
		gasFeeCap: gasFeeCap,
		gasTipCap: gasTipCap,
		data:      data,
		isFake:    isFake,
	}
}

//作为核心信息返回
func (tx *Transaction) AsMessage(s Signer, baseFee *big.Int) (Message, error) {
	msg := Message{
		nonce:     tx.Nonce(),
		gasLimit:  tx.Gas(),
		gasPrice:  new(big.Int).Set(tx.GasPrice()),
		gasFeeCap: new(big.Int).Set(tx.GasFeeCap()),
		gasTipCap: new(big.Int).Set(tx.GasTipCap()),
		to:        tx.To(),
		amount:    tx.Value(),
		data:      tx.Data(),
		isFake:    false,
	}
	//如果提供了baseFee，请将gasPrice设置为effectiveGasPrice。
	//if baseFee != nil {
	//	msg.gasPrice = BigMin(msg.gasPrice.Add(msg.gasTipCap, baseFee), msg.gasFeeCap)
	//}
	var err error
	msg.from, err = Sender(s, tx)
	return msg, err
}

func (m Message) From() utils.Address { return m.from }
func (m Message) To() *utils.Address  { return m.to }
func (m Message) GasPrice() *big.Int  { return m.gasPrice }
func (m Message) GasFeeCap() *big.Int { return m.gasFeeCap }
func (m Message) GasTipCap() *big.Int { return m.gasTipCap }
func (m Message) Value() *big.Int     { return m.amount }
func (m Message) Gas() uint64         { return m.gasLimit }
func (m Message) Nonce() uint64       { return m.nonce }
func (m Message) Data() []byte        { return m.data }
func (m Message) IsFake() bool        { return m.isFake }

func copyAddressPtr(a *utils.Address) *utils.Address {
	if a == nil {
		return nil
	}
	cpy := *a
	return &cpy
}
