package core

import (
	"sync"
)

var octopusCore *OctopusCore

var once sync.Once

//单例模式
func getInstance() *OctopusCore {
	once.Do(func() {
		octopusCore = new(OctopusCore)
	})
	return octopusCore
}

//创建octopusCore
func Create() {
	getInstance().create()
}

//依赖注入
func DependenceInject(Config map[interface{}]interface{}) {
	getInstance().setCfg(Config)
	getInstance().dependenceInject()
}

//依赖注入
func GetLang(langName string) interface{} {
	return getInstance().getLang(langName)
}

//销毁
func destroy() {
	getInstance().destroy()
}

//lang注册到core里面
func Register(lang interface{}) {
	getInstance().register(lang)
}

//lang注销core里面的实例
func Unregister(lang interface{}) {
	getInstance().unregister(lang)
}

func GetStartOctopusLang() map[string][]string {
	return getInstance().longsByStartMethod()
}

func GetStopOctopusLang() map[string][]string {
	return getInstance().longsByStopMethod()
}

func CallStartByName(name string) {
	getInstance().CallMethod(name, startMethodName, nil)
}

func CallStopByName(name string) {
	getInstance().CallMethod(name, stopMethodName, nil)
}

func CallMethod(name string, methodName string, in ...interface{}) []interface{} {
	return getInstance().CallMethod(name, methodName, in)
}
