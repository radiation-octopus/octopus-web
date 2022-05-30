package db

import "octopus/log"

//P2p启动方法
type P2pStart struct {
}

func (l *P2pStart) Start() {
	Start()
	log.Info("P2pStart start")
}
