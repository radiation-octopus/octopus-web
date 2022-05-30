package cache

import "octopus/log"

//Cache启动方法
type CacheStart struct {
	IsAutoCreate       bool `autoInjectCfg:"octopus.cache.auto.create.is"`
	AutoCreateCacheNum int  `autoInjectCfg:"octopus.cache.auto.create.num"`
}

func (c *CacheStart) Start() {
	IsAutoCreateCache = c.IsAutoCreate
	AutoCreateCacheNum = c.AutoCreateCacheNum
	Start()
	log.Info("CacheStart start")
}
