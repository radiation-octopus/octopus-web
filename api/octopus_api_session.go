package api

import "time"

type OctopusApiSession struct {
	Key          string
	value        map[string]interface{}
	ExpirateTime time.Time
}

func (s *OctopusApiSession) getSessionVal(sessionKey string) interface{} {
	return s.value[sessionKey]
}

func (s *OctopusApiSession) setSessionVal(sessionKey string, value interface{}) {
	s.value[sessionKey] = value
}

func initSession(key string) *OctopusApiSession {
	session := new(OctopusApiSession)
	session.Key = key
	session.value = make(map[string]interface{})
	session.ExpirateTime = time.Now()
	return session
}
