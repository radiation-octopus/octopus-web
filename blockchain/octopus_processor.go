package blockchain

import (
	"fmt"
	"math/big"
	"octopus/block"
	"octopus/consensus"
	"octopus/db"
	"octopus/log"
	"octopus/transition"
	"octopus/utils"
	"octopus/vm"
)

//处理器结构体
type BlockProcessor struct {
	//config *params.ChainConfig // 链配置
	bc     *BlockChain      // 标准链
	engine consensus.Engine // 共识引擎
}

//构建处理器
func NewBlockProcessor(bc *BlockChain, engine consensus.Engine) *BlockProcessor {
	bp := &BlockProcessor{
		bc:     bc,
		engine: engine,
	}
	return bp
}

//处理器接口
type Processor interface {
	//处理改变区块状态，将区块加入主链
	Process(block *block.Block, operationdb *db.OperationDB, cfg vm.Config) (block.Receipts, []*log.OctopusLog, uint64, error)
}

func (p *BlockProcessor) Process(b *block.Block, operationdb *db.OperationDB, cfg vm.Config) (block.Receipts, []*log.OctopusLog, uint64, error) {
	var (
		receipts    block.Receipts
		usedGas     = new(uint64)
		header      = b.Header()
		blockHash   = b.Hash()
		blockNumber = b.Number()
		allLogs     []*log.OctopusLog
		gp          = new(transition.GasPool).AddGas(b.GasLimit())
	)
	blockContext := vm.NewOVMBlockContext(header, p.bc, nil)
	//初始化虚拟机
	vmonv := vm.NewOVM(blockContext, vm.TxContext{}, operationdb, cfg)
	for i, tx := range b.Transactions() {
		msg, err := tx.AsMessage(block.MakeSigner(header.Number), header.BaseFee)
		if err != nil {
			return nil, nil, 0, fmt.Errorf("could not apply tx %d [%v]: %w", i, tx.Hash().Hex(), err)
		}
		receipt, err := applyTransaction(msg, gp, operationdb, blockNumber, blockHash, tx, usedGas, vmonv)
		if err != nil {
			return nil, nil, 0, fmt.Errorf("could not apply tx %d [%v]: %w", i, tx.Hash().Hex(), err)
		}
		receipts = append(receipts, receipt)
		allLogs = append(allLogs, receipt.Logs...)
	}

	return receipts, allLogs, *usedGas, nil
}

//处理事务
func applyTransaction(msg block.Message, gp *transition.GasPool, operationdb *db.OperationDB, blockNumber *big.Int, blockHash utils.Hash, tx *block.Transaction, usedGas *uint64, ovm *vm.OVM) (*block.Receipt, error) {

	//将事务应用于当前状态（包含在env中）。
	result, err := transition.ApplyMessage(msg, gp)
	if err != nil {
		return nil, err
	}

	*usedGas += result.UsedGas
	// Create a new receipt for the transaction, storing the intermediate root and gas used
	// by the tx.
	receipt := &block.Receipt{}
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = result.UsedGas
	receipt.BlockHash = blockHash
	receipt.BlockNumber = blockNumber
	return receipt, err
}
