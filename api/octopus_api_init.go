package api

import "octopus/director"

//初始化octopus
func init() {
	//把web启动注入
	director.Register(new(WebStart))
	//把web停止注入
	director.Register(new(WebStop))
}
