package blockchain

import (
	"octopus/block"
	"octopus/consensus"
)

type BlockValidator struct {
	//config *params.chainConfig 	//链配置
	bc     *BlockChain      //标准连
	engine consensus.Engine //共识引擎
}

//
func NewBlockValidator(blockchain *BlockChain, engine consensus.Engine) *BlockValidator {
	validator := &BlockValidator{
		bc:     blockchain,
		engine: engine,
	}
	return validator
}

//定义验证器处理接口
type Validator interface {
	//验证给定块内容
	validateBody(block *block.Block) error
}

//区块验证具体实现
func (v *BlockValidator) validateBody(block *block.Block) error {

	return nil
}
