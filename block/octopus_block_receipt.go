package block

import (
	"math/big"
	"octopus/log"
	"octopus/utils"
)

//收据代表交易的返回结果
type Receipt struct {
	TxHash  utils.Hash
	GasUsed uint64
	Logs    []*log.OctopusLog

	BlockHash        utils.Hash
	BlockNumber      *big.Int
	TransactionIndex uint
}

//收据列表
type Receipts []*Receipt
