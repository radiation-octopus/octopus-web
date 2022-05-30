package db

import "octopus/log"

//P2p停止方法
type P2pStop struct {
}

func (l *P2pStart) Stop() {
	Stop()
	log.Info("P2pStop stop")
}
