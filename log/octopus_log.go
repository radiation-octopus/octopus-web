package log

import (
	"fmt"
	"octopus/utils"
	"os"
	"time"
)

type LogMsg struct {
	LogLevel string
	Time     time.Time
	Contents string
	FuncInfo string
}

type OctopusLog struct {
	LogMsgChan chan *LogMsg
}

func (l *OctopusLog) start() {
	os.MkdirAll(LogSavePath, os.ModePerm)
	l.cleanLogTask()
	l.consume()
}

func (l *OctopusLog) cleanLogTask() {
	go func() {
		if !LogIsSave {
			return
		}
		for {
			//定时任务
			timenow := time.Now()
			var timeBefore time.Time
			switch LogSaveCutTime {
			case DayLevel:
				timeBefore = timenow.AddDate(0, 0, -LogSaveCutMax)
			case MonthsLevel:
				timeBefore = timenow.AddDate(0, -LogSaveCutMax, 0)
			case YearLevel:
				timeBefore = timenow.AddDate(-LogSaveCutMax, 0, 0)
			default:
				timeBefore = timenow.AddDate(0, 0, -LogSaveCutMax)
			}
			timeBeforeFormat := timeBefore.Format("20060102")
			var beforePath string
			if LogSaveIsCut {
				beforePath = LogSavePath + "/" + timeBeforeFormat
			} else {
				beforePath = LogSavePath + "/"
			}
			os.RemoveAll(beforePath)
			time := time.NewTimer(time.Hour * 24)
			<-time.C
		}
	}()
}

func (l *OctopusLog) consume() {
	go func() {
		for {
			select {
			case LogMsg := <-l.LogMsgChan:
				consoleLog(LogMsg)
				saveLog(LogMsg)
			}
		}
	}()
}

func (l *OctopusLog) stop() {
	close(l.LogMsgChan)
}

//判断是否控制台输出
func isConsoleLog(logMsg *LogMsg) bool {
	switch LogConsoleLevel {
	case DebugLevel:
		return true
	case InfoLevel:
		if logMsg.LogLevel == DebugLevel {
			return false
		}
	case WarnLevel:
		if logMsg.LogLevel == DebugLevel || logMsg.LogLevel == InfoLevel {
			return false
		}
	case ErrorLevel:
		if logMsg.LogLevel == DebugLevel || logMsg.LogLevel == InfoLevel || logMsg.LogLevel == WarnLevel {
			return false
		}
	}
	return true
}

//控制台输出
func consoleLog(logMsg *LogMsg) {
	if !isConsoleLog(logMsg) {
		return
	}
	switch logMsg.LogLevel {
	case DebugLevel:
		timeFormat := logMsg.Time.Format(LogTimeFormat)
		level := utils.GetColorStr(LogDebugColor, true, logMsg.LogLevel)
		funcInfo := utils.GetColorStr(LogDebugColor, true, logMsg.FuncInfo)
		contents := utils.GetColorStr(LogDebugColor, false, logMsg.Contents)
		str := timeFormat + "  " + level + "  " + funcInfo + "  " + contents
		fmt.Println(str)
	case InfoLevel:
		timeFormat := logMsg.Time.Format(LogTimeFormat)
		level := utils.GetColorStr(LogInfoColor, true, logMsg.LogLevel) + "  "
		funcInfo := utils.GetColorStr(LogInfoColor, true, logMsg.FuncInfo)
		contents := utils.GetColorStr(LogInfoColor, false, logMsg.Contents)
		str := timeFormat + "  " + level + "  " + funcInfo + "  " + contents
		fmt.Println(str)
	case WarnLevel:
		timeFormat := logMsg.Time.Format(LogTimeFormat)
		level := utils.GetColorStr(LogDebugColor, true, logMsg.LogLevel) + "  "
		funcInfo := utils.GetColorStr(LogWarnColor, true, logMsg.FuncInfo)
		contents := utils.GetColorStr(LogWarnColor, false, logMsg.Contents)
		str := timeFormat + "  " + level + "  " + funcInfo + "  " + contents
		fmt.Println(str)
	case ErrorLevel:
		timeFormat := logMsg.Time.Format(LogTimeFormat)
		level := utils.GetColorStr(LogDebugColor, true, logMsg.LogLevel)
		funcInfo := utils.GetColorStr(LogErrorColor, true, logMsg.FuncInfo)
		contents := utils.GetColorStr(LogErrorColor, false, logMsg.Contents)
		str := timeFormat + "  " + level + "  " + funcInfo + "  " + contents
		fmt.Println(str)
	}
}

//save保存
func saveLog(logMsg *LogMsg) {
	if !LogIsSave {
		return
	}
	switch logMsg.LogLevel {
	case DebugLevel:
		slaveTimeFormat := logMsg.Time.Format("20060102")
		writeStr := logMsg.Time.Format(LogTimeFormat) + "  " + logMsg.LogLevel + "  " + logMsg.FuncInfo + "  " + logMsg.Contents
		writeLog(slaveTimeFormat, LogDebugFile, writeStr)
	case InfoLevel:
		slaveTimeFormat := logMsg.Time.Format("20060102")
		writeStr := logMsg.Time.Format(LogTimeFormat) + "  " + logMsg.LogLevel + "  " + logMsg.FuncInfo + "  " + logMsg.Contents
		writeLog(slaveTimeFormat, LogInfoFile, writeStr)
	case WarnLevel:
		slaveTimeFormat := logMsg.Time.Format("20060102")
		writeStr := logMsg.Time.Format(LogTimeFormat) + "  " + logMsg.LogLevel + "  " + logMsg.FuncInfo + "  " + logMsg.Contents
		writeLog(slaveTimeFormat, LogWarnFile, writeStr)
	case ErrorLevel:
		slaveTimeFormat := logMsg.Time.Format("20060102")
		writeStr := logMsg.Time.Format(LogTimeFormat) + "  " + logMsg.LogLevel + "  " + logMsg.FuncInfo + "  " + logMsg.Contents
		writeLog(slaveTimeFormat, LogErrorFile, writeStr)
	}
}

func writeLog(cutPath string, LogLevelFile string, writeStr string) {
	var path string
	if LogSaveIsCut {
		path = LogSavePath + "/" + cutPath + "/" + LogLevelFile
	} else {
		path = LogSavePath + "/" + LogLevelFile
	}
	utils.WriteFileLine(path, writeStr)
}
