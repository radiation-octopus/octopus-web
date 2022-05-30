package cache

import "octopus/director"

//初始化octopus
func init() {
	//把启动注入
	director.Register(new(CacheStart))
	//把停止注入
	director.Register(new(CacheStop))
}
