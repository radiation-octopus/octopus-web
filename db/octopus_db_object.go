package db

import (
	"math/big"
	"octopus/utils"
)

type Database struct {
}

//账户结构体
type StateAccount struct {
	Nonce    uint64
	Balance  *big.Int
	Root     utils.Hash // merkle root of the storage trie
	CodeHash []byte
}

//数据库存储对象
type OperationObject struct {
	address  utils.Address
	addrHash utils.Hash
	data     StateAccount
	db       *OperationDB
}
