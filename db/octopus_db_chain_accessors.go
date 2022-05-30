package db

import (
	"errors"
	"math/big"
	"octopus/block"
	"octopus/log"
	"octopus/utils"
)

//定义读取数据所需的方法
type Reader interface {
	KeyValueReader
	AncientReader
}

type KeyValueReader interface {
	//是否存在键
	IsHas(mark string, key string) (bool, error)

	//通过键检索值
	Query(mark string, key string) *block.Header
}

type KeyValueWriter interface {
	// Put inserts the given value into the key-value data store.
	Put(mark string, key string, value *big.Int) error

	// Delete removes the key from the key-value data store.
	Delete(mark string, key string)
}

func (db Database) IsHas(mark string, key string) (bool, error) {
	return IsHas(mark, key)
}
func (db Database) Query(mark string, key string) *block.Header {
	u := block.Header{}
	u = Query(mark, key, u).(block.Header)
	return &u
}

func (db Database) Put(mark string, key string, value *big.Int) error {
	return Insert(mark, key, value)
}

func (db Database) Delete(mark string, key string) {
	Delete(mark, key)
}

type AncientReader interface {
	//AncientReaderOp
	//
	//// ReadAncients runs the given read operation while ensuring that no writes take place
	//// on the underlying freezer.
	//ReadAncients(fn func(AncientReaderOp) error) (err error)
}

//检索与哈希对应的块头。
func ReadHeader(db Reader, hash utils.Hash, number uint64) *block.Header {
	return db.Query(headerMark, utils.GetInToStr(hash))
}

//td新增到数据库
func WriteTd(db KeyValueWriter, hash utils.Hash, number uint64, td *big.Int) {
	//data, err := rlp.EncodeToBytes(td)
	//if err != nil {
	//	log.Crit("Failed to RLP encode block total difficulty", "err", err)
	//}
	if err := db.Put(tdMark, utils.GetInToStr(hash), td); err != nil {
		errors.New("Failed to store block total difficulty")
	}
}
func WriteBlock(block *block.Block) {
	WriteBody(block.Hash(), block.Body())
	WriteHeader(block.Header())
}

//body新增到数据库
func WriteBody(hash utils.Hash, body *block.Body) {
	if err := Insert(tdMark, utils.GetInToStr(hash), body); err != nil {
		log.Info("Failed to store block total difficulty")
	}
}

//header新增到数据库
func WriteHeader(header *block.Header) {
	var (
		hash = header.Hash()
	)
	if err := Insert(headerMark, utils.GetInToStr(hash), header); err != nil {
		log.Info("Failed to store header")
	}
}

//收据新增到数据库
func WriteReceipts(hash utils.Hash, receipts block.Receipts) {
	if err := Insert(receiptsMark, utils.GetInToStr(hash), receipts); err != nil {
		log.Info("Failed to store header")
	}
}
