package db

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"octopus/utils"
	"sync"
)

type DatabaseI interface {
	// 打开主帐户trie。
	OpenTrie(root utils.Hash) (Trie, error)

	// 打开帐户的存储trie。
	OpenStorageTrie(addrHash, root utils.Hash) (Trie, error)

	// CopyTrie returns an independent copy of the given trie.
	CopyTrie(Trie) Trie

	// ContractCode retrieves a particular contract's code.
	ContractCode(addrHash, codeHash utils.Hash) ([]byte, error)

	// ContractCodeSize retrieves a particular contracts code's size.
	ContractCodeSize(addrHash, codeHash utils.Hash) (int, error)

	// TrieDB retrieves the low level trie database used for data storage.
	TrieDB() *Database
}

type Trie interface {
	// GetKey returns the sha3 preimage of a hashed key that was previously used
	// to store a value.
	//
	// TODO(fjl): remove this when SecureTrie is removed
	GetKey([]byte) []byte

	// TryGet returns the value for key stored in the trie. The value bytes must
	// not be modified by the caller. If a node was not found in the database, a
	// trie.MissingNodeError is returned.
	TryGet(key []byte) ([]byte, error)

	// TryUpdateAccount abstract an account write in the trie.
	TryUpdateAccount(key []byte, account *StateAccount) error

	// TryUpdate associates key with value in the trie. If value has length zero, any
	// existing value is deleted from the trie. The value bytes must not be modified
	// by the caller while they are stored in the trie. If a node was not found in the
	// database, a trie.MissingNodeError is returned.
	TryUpdate(key, value []byte) error

	// TryDelete removes any existing value for key from the trie. If a node was not
	// found in the database, a trie.MissingNodeError is returned.
	TryDelete(key []byte) error

	// Hash returns the root hash of the trie. It does not write to the database and
	// can be used even if the trie doesn't have one.
	Hash() utils.Hash

	// Commit writes all nodes to the trie's memory database, tracking the internal
	// and external (for account tries) references.
	//Commit(onleaf trie.LeafCallback) (blockchain.Hash, int, error)

	// NodeIterator returns an iterator that returns nodes of the trie. Iteration
	// starts at the key after the given start key.
	//NodeIterator(startKey []byte) trie.NodeIterator

	// Prove constructs a Merkle proof for key. The result contains all encoded nodes
	// on the path to the value at key. The value itself is also included in the last
	// node and can be retrieved by verifying the proof.
	//
	// If the trie does not contain a value for key, the returned proof contains all
	// nodes of the longest existing prefix of the key (at least the root), ending
	// with the node that proves the absence of the key.
	//Prove(key []byte, fromLevel uint, proofDb KeyValueWriter) error
}

var db *leveldb.DB

var once sync.Once

//单例模式
func getInstance() *leveldb.DB {
	once.Do(func() {
		db, _ = leveldb.OpenFile(SaveDbFilePath, nil)
	})
	return db
}

func Start() {
	getInstance()
}

func Stop() {
	getInstance().Close()
}

func getKeyStr(mark string, key string) string {
	return mark + "-" + key
}

//插入
func Insert(mark string, key string, value interface{}) error {
	keyStr := getKeyStr(mark, key)
	valuebyte, error := json.Marshal(&value)
	getInstance().Put([]byte(keyStr), []byte(valuebyte), nil)
	if error != nil {
		return error
	} else {
		return nil
	}
}

//删除
func Delete(mark string, key string) {
	keyStr := getKeyStr(mark, key)
	getInstance().Delete([]byte(keyStr), nil)
}

//查询
func Query(mark string, key string, value interface{}) interface{} {
	keyStr := getKeyStr(mark, key)
	//valueStr := utils.GetInToStr(value)
	valueByte, _ := getInstance().Get([]byte(keyStr), nil)
	json.Unmarshal(valueByte, &value)
	return value
}

//读取mark的所有数据
func QueryMark(mark string) map[string]string {
	iter := db.NewIterator(util.BytesPrefix([]byte(mark+"-")), nil)
	var returnMap map[string]string
	for iter.Next() {
		key := string(iter.Key())
		value := string(iter.Value())
		returnMap[key] = value
	}
	return returnMap
}

//是否包含
func IsHas(mark string, key string) (bool, error) {
	keyStr := getKeyStr(mark, key)
	//valueStr := utils.GetInToStr(value)
	bl, error := getInstance().Has([]byte(keyStr), nil)
	return bl, error
}

//查询给定键
func (db *Database) Get(key []byte) ([]byte, error) {
	dat, err := getInstance().Get(key, nil)
	if err != nil {
		return nil, err
	}
	return dat, nil
}
