package tcp

import (
	"octopus/log"
)

//Tcp停止方法
type TcpStop struct {
}

func (t *TcpStart) Stop() {
	Stop()
	log.Info("TcpStop stop")
}
