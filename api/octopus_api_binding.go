package api

import (
	"sync"
)

type OctopusApiBinding struct {
	CallFunc func(map[string]interface{}) map[string]interface{}
	BaseApi  BaseApi
	Mutex    *sync.Mutex
}

func initApiBinding(baseApi BaseApi, f func(map[string]interface{}) map[string]interface{}) *OctopusApiBinding {
	apiBinding := new(OctopusApiBinding)
	apiBinding.BaseApi = baseApi
	apiBinding.CallFunc = f
	apiBinding.Mutex = getInstance().getMutex(baseApi)
	return apiBinding
}
