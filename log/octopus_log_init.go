package log

import "octopus/director"

//初始化octopus
func init() {
	log := getInstance()
	log.LogMsgChan = make(chan *LogMsg, LogMsgNum)
	//把启动注入
	director.Register(new(LogStart))
	//把停止注入
	director.Register(new(LogStop))
}
