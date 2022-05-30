package udp

import "octopus/log"

//Udp启动方法
type UdpStart struct {
	Port           int    `autoInjectCfg:"octopus.udp.port"`
	MsgNum         int    `autoInjectCfg:"octopus.udp.msg.num"`
	BindingPoolNum int    `autoInjectCfg:"octopus.udp.binding.pool.num"`
	BindingMethod  string `autoInjectCfg:"octopus.udp.binding.method"`
	BindingStruct  string `autoInjectCfg:"octopus.udp.binding.struct"`
}

func (u *UdpStart) Start() {
	UdpPort = u.Port
	UdpMsgNum = u.MsgNum
	UdpAcceptCallBindingMethod = u.BindingMethod
	UdpAcceptCallBindingStruct = "*" + u.BindingStruct
	UdpAcceptCallBindingPoolNum = u.BindingPoolNum
	Start()
	log.Info("UdpStart start")
}
