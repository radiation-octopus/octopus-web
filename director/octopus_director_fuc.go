package director

import (
	"octopus/core"
	"octopus/utils"
	"sync"
)

var octopusDirector *OctopusDirector

var once sync.Once

//单例模式
func getInstance() *OctopusDirector {
	once.Do(func() {
		octopusDirector = new(OctopusDirector)
	})
	return octopusDirector
}

//启动
func Start() {
	getInstance().directorStart()
}

//停止
func Stop() {
	getInstance().directorStop()
}

//配置文件路径
func SetProfilesCfgPath(cfgLocationPath string) {
	getInstance().setProfilesCfgPath(cfgLocationPath)
}

////配置文件前缀
//func SetProfilesCfgPrefix(profilesCfgPrefix string) {
//	getInstance().setProfilesCfgPrefix(profilesCfgPrefix)
//}
//
////配置文件后缀
//func SetProfilesCfgType(profilesCfgType string) {
//	getInstance().setProfilesCfgType(profilesCfgType)
//}

func loadProfilesCfgPathes(profilesCfgPathes []string, profilesCfgPrefix string, profilesCfgType string) map[interface{}]interface{} {
	var cfg map[interface{}]interface{}
	for _, profilesCfgPath := range profilesCfgPathes {
		cfg = loadProfilesCfgPath(profilesCfgPath, profilesCfgPrefix, profilesCfgType, "", cfg)
	}
	return cfg
}

//读配置文件放入Cfg里
func loadProfilesCfgPath(profilesCfgPath string, profilesCfgPrefix string, profilesCfgType string, activeName string, cfg map[interface{}]interface{}) map[interface{}]interface{} {
	getcfg := map[interface{}]interface{}{}
	if activeName == "" {
		path := profilesCfgPath + "/" + profilesCfgPrefix + "." + profilesCfgType
		getcfg = utils.ReadYaml(path)
		cfg = utils.CompareMap(cfg, getcfg)
	} else {
		path := profilesCfgPath + "/" + profilesCfgPrefix + "-" + activeName + "." + profilesCfgType
		getcfg = utils.ReadYaml(path)
		cfg = utils.CompareMap(cfg, getcfg)
	}
	octopus := getcfg[ProfilesCfgDefaultKey1].(map[interface{}]interface{})
	if octopus == nil {
		return cfg
	}
	octopusCfg := octopus[ProfilesCfgDefaultKey2]
	if octopusCfg != nil {
		active := octopusCfg.(map[interface{}]interface{})[ProfilesCfgDefaultKey3]
		if active != nil {
			activeName = active.(string)
			cfg = loadProfilesCfgPath(profilesCfgPath, profilesCfgPrefix, profilesCfgType, activeName, cfg)
		}
	}
	return cfg
}

//读配置文件变量
func ReadCfg(path ...string) interface{} {
	return getInstance().readCfg(path)
}

//修改配置文件变量
func WriteCfg(writeCfg interface{}, path ...string) {
	getInstance().writeCfg(writeCfg, path)
}

//lang注册到core里面
func Register(lang interface{}) {
	core.Register(lang)
}

//lang注销core里面的实例
func Unregister(lang interface{}) {
	core.Unregister(lang)
}

func FindHeadOrders(deepMap map[string][]string) []string {
	order := []string{}
	for k := range deepMap {
		order = findDeepOrder(deepMap, k, order)
	}
	newOrder := []string{}
	for i := range order {
		newOrder = append(newOrder, order[len(order)-(i+1)])
	}
	return newOrder
}

func FindDeepOrders(deepMap map[string][]string) []string {
	order := []string{}
	for k := range deepMap {
		order = findDeepOrder(deepMap, k, order)
	}
	return order
}

func findDeepOrder(deepMap map[string][]string, key string, order []string) []string {
	value := deepMap[key]
	if value != nil && len(value) != 0 {
		for i := range value {
			order = findDeepOrder(deepMap, value[i], order)
		}
		bl := true
		for i := range order {
			if key == order[i] {
				bl = false
				break
			}
		}
		if bl {
			order = append(order, key)
		}
	} else {
		bl := true
		for i := range order {
			if key == order[i] {
				bl = false
				break
			}
		}
		if bl {
			order = append(order, key)
		}
	}
	return order
}
