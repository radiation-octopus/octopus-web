package log

//Log停止方法
type LogStop struct {
}

func (l *LogStart) Stop() {
	Stop()
	Info("LogStop stop")
}
