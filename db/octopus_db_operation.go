package db

import (
	"math/big"
	"octopus/log"
	"octopus/utils"
)

//数据库操作结构体
type OperationDB struct {
	db           Database
	originalRoot utils.Hash //根hash
	//trie         			Trie
	OperationObjects map[utils.Address]*OperationObject // 数据库活动对象缓存map集合
	dbErr            error                              // 数据库操作错误记录存储

	thash   utils.Hash
	txIndex int
	logs    map[utils.Hash][]*log.OctopusLog //日志
	logSize uint

	preimages map[utils.Hash][]byte

	accessList *accessList // 事务访问进程列表

	//AccountReads         time.Duration
	//AccountHashes        time.Duration
	//AccountUpdates       time.Duration
	//AccountCommits       time.Duration
	//StorageReads         time.Duration
	//StorageHashes        time.Duration
	//StorageUpdates       time.Duration
	//StorageCommits       time.Duration
	//SnapshotAccountReads time.Duration
	//SnapshotStorageReads time.Duration
	//SnapshotCommits      time.Duration
	//
	//AccountUpdated int
	//StorageUpdated int
	//AccountDeleted int
	//StorageDeleted int
}

func (o OperationDB) CreateAccount(address utils.Address) {

	panic("implement me")
}

func (o OperationDB) SubBalance(address utils.Address, b *big.Int) {
	panic("implement me")
}

func (o OperationDB) AddBalance(address utils.Address, b *big.Int) {
	panic("implement me")
}

func (o OperationDB) GetBalance(address utils.Address) *big.Int {
	panic("implement me")
}

func (o OperationDB) GetNonce(address utils.Address) uint64 {
	panic("implement me")
}

func (o OperationDB) SetNonce(address utils.Address, u uint64) {
	panic("implement me")
}

func (o OperationDB) GetCodeHash(address utils.Address) utils.Hash {
	panic("implement me")
}

func (o OperationDB) GetCode(address utils.Address) []byte {
	panic("implement me")
}

func (o OperationDB) SetCode(address utils.Address, bytes []byte) {
	panic("implement me")
}

func (o OperationDB) GetCodeSize(address utils.Address) int {
	panic("implement me")
}

func (o OperationDB) AddRefund(u uint64) {
	panic("implement me")
}

func (o OperationDB) SubRefund(u uint64) {
	panic("implement me")
}

func (o OperationDB) GetRefund() uint64 {
	panic("implement me")
}

func (o OperationDB) GetCommittedState(address utils.Address, hash utils.Hash) utils.Hash {
	panic("implement me")
}

func (o OperationDB) GetState(address utils.Address, hash utils.Hash) utils.Hash {
	panic("implement me")
}

func (o OperationDB) SetState(address utils.Address, hash utils.Hash, hash2 utils.Hash) {
	panic("implement me")
}

func (o OperationDB) Suicide(address utils.Address) bool {
	panic("implement me")
}

func (o OperationDB) HasSuicided(address utils.Address) bool {
	panic("implement me")
}

func (o OperationDB) Exist(address utils.Address) bool {
	panic("implement me")
}

func (o OperationDB) Empty(address utils.Address) bool {
	panic("implement me")
}

func (o OperationDB) AddressInAccessList(addr utils.Address) bool {
	panic("implement me")
}

func (o OperationDB) SlotInAccessList(addr utils.Address, slot utils.Hash) (addressOk bool, slotOk bool) {
	panic("implement me")
}

func (o OperationDB) AddAddressToAccessList(addr utils.Address) {
	panic("implement me")
}

func (o OperationDB) AddSlotToAccessList(addr utils.Address, slot utils.Hash) {
	panic("implement me")
}

func (o OperationDB) RevertToSnapshot(i int) {
	panic("implement me")
}

func (o OperationDB) Snapshot() int {
	panic("implement me")
}

func (o OperationDB) AddLog(octopusLog *log.OctopusLog) {
	panic("implement me")
}

func (o OperationDB) AddPreimage(hash utils.Hash, bytes []byte) {
	panic("implement me")
}

func (o OperationDB) ForEachStorage(address utils.Address, f func(utils.Hash, utils.Hash) bool) error {
	panic("implement me")
}

//创建根状态数据库
func New(root utils.Hash, db Database) (*OperationDB, error) {
	//tr, err := db.OpenTrie(root)
	//if err != nil {
	//	return nil, err
	//}
	sdb := &OperationDB{
		db: db,
		//trie:                tr,
		originalRoot:     root,
		OperationObjects: make(map[utils.Address]*OperationObject),
		logs:             make(map[utils.Hash][]*log.OctopusLog),
		preimages:        make(map[utils.Hash][]byte),
		accessList:       newAccessList(),
	}
	return sdb, nil
}

/**
accessList
*/
type accessList struct {
	addresses map[utils.Address]int
	slots     []map[utils.Hash]struct{}
}

//创建集合
func newAccessList() *accessList {
	return &accessList{
		addresses: make(map[utils.Address]int),
	}
}
