package log

//存放默认常量

//是否保存日志
var LogIsSave = true

//日志存储地址
var LogSavePath = "logs"

//保存分段 （年，月，日）
var LogSaveIsCut = true

//保存分段 （年，月，日）
var LogSaveCutTime = "day"

//日志最长保存时间
var LogSaveCutMax = 30

//控制台输出等级
var LogConsoleLevel = "debug"

//是否保存日志
var LogTimeFormat = "2001-05-16 15:04:05 500"

//日志数量
var LogMsgNum = 1024

//debug 颜色
var LogDebugColor = "cyan"

//debug 存储文件名称
var LogDebugFile = "debug.log"

//info 颜色
var LogInfoColor = "blue"

//info 存储文件名称
var LogInfoFile = "info.log"

//warn 颜色
var LogWarnColor = "yellow"

//warn 存储文件名称
var LogWarnFile = "warn.log"

//error 颜色
var LogErrorColor = "red"

//error 存储文件名称
var LogErrorFile = "error.log"
