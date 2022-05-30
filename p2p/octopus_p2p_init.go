package db

import "octopus/director"

//初始化octopus
func init() {
	//把启动注入
	director.Register(new(P2pStart))
	//把停止注入
	director.Register(new(P2pStop))
	//把udp回调注入
	director.Register(new(P2pCallUdp))
	//把tcp Clinet回调注入
	director.Register(new(P2pCallTcpClinet))
	//把tcp Serve回调注入
	director.Register(new(P2pCallTcpServer))
}
