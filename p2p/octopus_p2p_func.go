package db

import (
	"sync"
)

var octopusP2p *OctopusP2p

var once sync.Once

//单例模式
func getInstance() *OctopusP2p {
	once.Do(func() {
		octopusP2p = new(OctopusP2p)
	})
	return octopusP2p
}

func Start() {
	getInstance().start()
}

func Stop() {
	getInstance().stop()
}
