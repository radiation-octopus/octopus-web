package tcp

import "octopus/log"

//Tcp启动方法
type TcpStart struct {
	ServerPort           int    `autoInjectCfg:"octopus.tcp.server.port"`
	ServerBindingPoolNum int    `autoInjectCfg:"octopus.tcp.server.binding.pool.num"`
	ServerBindingMethod  string `autoInjectCfg:"octopus.tcp.server.binding.method"`
	ServerBindingStruct  string `autoInjectCfg:"octopus.tcp.server.binding.struct"`
	ClinetPort           int    `autoInjectCfg:"octopus.tcp.clinet.port"`
	ClinetMsgNum         int    `autoInjectCfg:"octopus.tcp.clinet.msg.num"`
	ClinetBindingMethod  string `autoInjectCfg:"octopus.tcp.clinet.binding.method"`
	ClinetBindingStruct  string `autoInjectCfg:"octopus.tcp.clinet.binding.struct"`
}

func (t *TcpStart) Start() {
	TcpClinetMsgNum = t.ClinetMsgNum
	TcpServerPort = t.ServerPort
	TcpServerAcceptCallBindingPoolNum = t.ServerBindingPoolNum
	TcpServerAcceptCallBindingMethod = t.ServerBindingMethod
	TcpServerAcceptCallBindingStruct = "*" + t.ServerBindingStruct
	TcpClinetPort = t.ClinetPort
	TcpClinetAcceptCallBindingMethod = t.ClinetBindingMethod
	TcpClinetAcceptCallBindingStruct = "*" + t.ClinetBindingStruct
	Start()
	log.Info("TcpStart start")
}
