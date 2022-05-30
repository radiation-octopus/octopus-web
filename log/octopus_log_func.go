package log

import (
	"octopus/utils"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var octopusLog *OctopusLog

var once sync.Once

//单例模式
func getInstance() *OctopusLog {
	once.Do(func() {
		octopusLog = new(OctopusLog)
	})
	return octopusLog
}

var mutex = &sync.Mutex{}

func CreateLogMsg(logLevel string, contents string) {
	logMsg := new(LogMsg)
	logMsg.Time = time.Now()
	logMsg.LogLevel = logLevel
	logMsg.Contents = contents
	mutex.Lock()
	_, file, lineNo, _ := runtime.Caller(2)
	mutex.Unlock()
	//funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // Base函数返回路径的最后一个元素
	logMsg.FuncInfo = fileName + ":" + strconv.Itoa(lineNo)
	getInstance().LogMsgChan <- logMsg
}

func Start() {
	getInstance().start()
}

func Stop() {
	getInstance().stop()
}

func Debug(in ...interface{}) {
	str := utils.GetInToStr(in)
	CreateLogMsg(DebugLevel, str)
}

func Info(in ...interface{}) {
	str := utils.GetInToStr(in)
	CreateLogMsg(InfoLevel, str)
}

func Warn(in ...interface{}) {
	str := utils.GetInToStr(in)
	CreateLogMsg(WarnLevel, str)
}

func Error(in ...interface{}) {
	str := utils.GetInToStr(in)
	CreateLogMsg(ErrorLevel, str)
}
