package api

//存放默认常量

//是否开启session
var IsSession bool

//session保存时间类型
var SessionTime string

//session 的最大存放时间
var SessionMaxAge int

//session cookie的路径
var SessionCookiePath string

//session cookie的站点地址
var SessionCookieDomain string

//是否开启cookie
var IsCookie bool

//cookie保存时间类型
var CookieTime string

//默认 cookie 地址
var CookiePath string

//默认 cookie 网站存储点
var CookieDomain string

//默认 cookie 存储时间
var CookieMaxAge int

//路径值前缀
var RouteMapPathVariablePrefix string

//路径方法前缀
var RouteMapMethodPrefix string

//http协议线程池最大数量
var ApiHandlePoolNum int
