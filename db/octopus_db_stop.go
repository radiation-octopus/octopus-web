package db

import "octopus/log"

//Db停止方法
type DbStop struct {
}

func (d *DbStart) Stop() {
	Stop()
	log.Info("DbStop stop")
}
