package blockchain

import "octopus/director"

//初始化octopus_blockchain
func init() {
	//把启动注入
	director.Register(new(BlockChainStart))
	//把停止注入
	director.Register(new(BlockChainStop))
}
