package udp

import "octopus/director"

//初始化octopus
func init() {
	//把启动注入
	director.Register(new(UdpStart))
	//把停止注入
	director.Register(new(UdpStop))
}
