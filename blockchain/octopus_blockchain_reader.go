package blockchain

import (
	"math/big"
	"octopus/block"
	"octopus/utils"
)

func (bc *BlockChain) getBlockByNumber(number uint64) *block.Block {
	//把创世区块注入
	//director.Register(new(Block))

	return nil
}

func (bc *BlockChain) CurrentHeader() *block.Header {
	return bc.hc.CurrentHeader()
}

func (bc *BlockChain) GetHeader(hash utils.Hash, number uint64) *block.Header {
	return bc.hc.GetHeader(hash, number)
}

func (bc *BlockChain) GetHeaderByHash(hash utils.Hash) *block.Header {
	return bc.hc.GetHeaderByHash(hash)
}

func (bc *BlockChain) GetHeaderByNumber(number uint64) *block.Header {
	return bc.hc.GetHeaderByNumber(number)
}

func (bc *BlockChain) GetTd(hash utils.Hash, number uint64) *big.Int {
	return bc.hc.GetTd(hash, number)
}
