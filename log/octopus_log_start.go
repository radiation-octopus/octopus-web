package log

//Log启动方法
type LogStart struct {
	IsSave      bool   `autoInjectCfg:"octopus.log.save.is"`
	SavePath    string `autoInjectCfg:"octopus.log.save.path"`
	SaveIsCut   bool   `autoInjectCfg:"octopus.log.save.cut.is"`
	SaveCutTime string `autoInjectCfg:"octopus.log.save.cut.time"`
	SaveCutMax  int    `autoInjectCfg:"octopus.log.save.cut.max"`

	ConsoleLevel string `autoInjectCfg:"octopus.log.console.level"`

	TimeFormat string `autoInjectCfg:"octopus.log.time.format"`

	MsgNum int `autoInjectCfg:"octopus.log.msg.num"`

	DebugColor string `autoInjectCfg:"octopus.log.debug.color"`
	DebugFile  string `autoInjectCfg:"octopus.log.debug.file"`

	InfoColor string `autoInjectCfg:"octopus.log.info.color"`
	InfoFile  string `autoInjectCfg:"octopus.log.info.file"`

	WarnColor string `autoInjectCfg:"octopus.log.warn.color"`
	WarnFile  string `autoInjectCfg:"octopus.log.warn.file"`

	ErrorColor string `autoInjectCfg:"octopus.log.error.color"`
	ErrorFile  string `autoInjectCfg:"octopus.log.error.file"`
}

func (l *LogStart) Start() {
	LogIsSave = l.IsSave
	LogSavePath = l.SavePath
	LogSaveIsCut = l.SaveIsCut
	LogSaveCutTime = l.SaveCutTime
	LogSaveCutMax = l.SaveCutMax

	LogConsoleLevel = l.ConsoleLevel

	LogTimeFormat = l.TimeFormat

	LogMsgNum = l.MsgNum

	LogDebugColor = l.DebugColor
	LogDebugFile = l.DebugFile

	LogInfoColor = l.InfoColor
	LogInfoFile = l.InfoFile

	LogWarnColor = l.WarnColor
	LogWarnFile = l.WarnFile

	LogErrorColor = l.ErrorColor
	LogErrorFile = l.ErrorFile

	Start()
	Info("LogStart start")
}
