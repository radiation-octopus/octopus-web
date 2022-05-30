package blockchain

import (
	lru "github.com/hashicorp/golang-lru"
	"math/big"
	mrand "math/rand"
	"octopus/block"
	"octopus/consensus"
	"octopus/db"
	"octopus/utils"
	"sync/atomic"
)

type HeaderChain struct {
	//config        *params.ChainConfig
	chainDb       db.Database
	genesisHeader *block.Header //创世区块的当前头部

	currentHeader     atomic.Value //头链的当前头部
	currentHeaderHash utils.Hash   //头链的当前头的hash

	headerCache *lru.Cache // 缓存最近的块头
	tdCache     *lru.Cache // 缓存最近的块总困难数
	numberCache *lru.Cache // 缓存最近的块号

	procInterrupt func() bool

	rand   *mrand.Rand
	engine consensus.Engine
}

func (hc *HeaderChain) CurrentHeader() *block.Header {
	return hc.currentHeader.Load().(*block.Header)
}
func (hc *HeaderChain) GetHeader(hash utils.Hash, number uint64) *block.Header {
	//先查看缓存通道是否存在
	if header, ok := hc.headerCache.Get(hash); ok {
		return header.(*block.Header)
	}
	//检索数据库查询
	header := db.ReadHeader(hc.chainDb, hash, number)
	if header == nil {
		return nil
	}
	// Cache the found header for next time and return
	hc.headerCache.Add(hash, header)
	return header
}
func (hc *HeaderChain) GetHeaderByNumber(number uint64) *block.Header {
	return nil
}
func (hc *HeaderChain) GetHeaderByHash(hash utils.Hash) *block.Header {
	return nil
}
func (hc *HeaderChain) GetTd(hash utils.Hash, number uint64) *big.Int {
	return nil
}
