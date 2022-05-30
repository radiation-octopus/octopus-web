package main

import (
	_ "octopus/api"
	_ "octopus/blockchain"
	"octopus/core"
	_ "octopus/db"
	"octopus/director"
	"octopus/log"
	_ "octopus/log"
	_ "octopus/tcp"
	_ "octopus/udp"
)

func init() {
	mainDao := new(MainDao)
	mainService := new(MainService)
	director.Register(mainDao)
	director.Register(mainService)
}

func main() {
	director.Start()
	var testMainDao = core.GetLang("main.MainDao").(*MainDao)
	var testMainService = core.GetLang("main.MainService").(*MainService)
	log.Info(testMainDao, testMainService)
}
