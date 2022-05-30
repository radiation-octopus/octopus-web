package db

import "octopus/log"

//Db启动方法
type DbStart struct {
	SavePath string `autoInjectCfg:"octopus.db.save.path"`
}

func (d *DbStart) Start() {
	SaveDbFilePath = d.SavePath
	Start()
	log.Info("DbStart start")

}
