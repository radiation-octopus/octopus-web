package consensus

import (
	"math/big"
	"octopus/block"
	"octopus/utils"
)

//该接口定义了验证期间访问本地本地区块两所需的一小部分方法
type ChainHeaderReader interface {
	// Config retrieves the blockchain's chain configuration.
	//Config() *params.ChainConfig

	// 从本地链检索当前头
	CurrentHeader() *block.Header

	// 通过hash和数字从数据库检索块头
	GetHeader(hash utils.Hash, number uint64) *block.Header

	// 按编号从数据库检索块头
	GetHeaderByNumber(number uint64) *block.Header

	// 通过其hash从数据库中检索块头
	GetHeaderByHash(hash utils.Hash) *block.Header

	// 通过hash和数字从数据库中检索总难度
	GetTd(hash utils.Hash, number uint64) *big.Int
}

//共识引擎接口
type Engine interface {
	Author(header *block.Header) (utils.Address, error)
	//表头验证器，该方法返回退出通道以终止操作，验证顺序为切片排序
	VerifyHeaders(chain ChainHeaderReader, headers []*block.Header, seals []bool) (chan<- struct{}, <-chan error)
}
