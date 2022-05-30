package cache

import "octopus/log"

//Cache停止方法
type CacheStop struct {
}

func (d *CacheStart) Stop() {
	Stop()
	log.Info("CacheStop stop")
}
