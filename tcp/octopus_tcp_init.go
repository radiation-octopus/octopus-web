package tcp

import "octopus/director"

//初始化octopus
func init() {
	//把启动注入
	director.Register(new(TcpStart))
	//把停止注入
	director.Register(new(TcpStop))
}
