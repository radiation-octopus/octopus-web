package api

import "octopus/log"

//Web停止方法
type WebStop struct {
}

func (w *WebStop) Stop() {
	Stop()
	log.Info("WebStop stop")
}
