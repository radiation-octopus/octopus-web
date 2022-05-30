package api

import (
	"octopus/log"
)

//WebStart
type WebStart struct {
	IsSession           bool   `autoInjectCfg:"octopus.api.session.is"`
	SessionMaxAge       int    `autoInjectCfg:"octopus.api.session.maxAge"`
	SessionCookiePath   string `autoInjectCfg:"octopus.api.session.cookie.path"`
	SessionCookieDomain string `autoInjectCfg:"octopus.api.session.cookie.website"`
	IsCookie            bool   `autoInjectCfg:"octopus.api.cookie.is"`
	CookiePath          string `autoInjectCfg:"octopus.api.cookie.path"`
	CookieDomain        string `autoInjectCfg:"octopus.api.cookie.website"`
	CookieMaxAge        int    `autoInjectCfg:"octopus.api.cookie.maxAge"`
	PathVariablePrefix  string `autoInjectCfg:"octopus.api.pathVariable.prefix"`
	MethodPrefix        string `autoInjectCfg:"octopus.api.method.prefix"`
	ApiHandlePoolNum    int    `autoInjectCfg:"octopus.api.handle.pool.num"`
}

func (w *WebStart) Start() {
	IsSession = w.IsSession
	SessionMaxAge = w.SessionMaxAge
	SessionCookiePath = w.SessionCookiePath
	SessionCookieDomain = w.SessionCookieDomain
	IsCookie = w.IsCookie
	CookiePath = w.CookiePath
	CookieDomain = w.CookieDomain
	CookieMaxAge = w.CookieMaxAge
	RouteMapPathVariablePrefix = w.PathVariablePrefix
	RouteMapMethodPrefix = w.MethodPrefix
	ApiHandlePoolNum = w.ApiHandlePoolNum
	Start()
	log.Info("ApiStart start")
	CleanSessionSchedule()
}
