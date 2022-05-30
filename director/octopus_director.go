package director

import (
	"octopus/core"
	"reflect"
	"strings"
)

type OctopusDirector struct {
	ProfilesCfgPath       []string                    //配置文件路劲
	ProfilesCfgType       string                      //配置文件类型
	ProfilesCfgPrefix     string                      //配置文件前缀
	Config                map[interface{}]interface{} //配置文件详情
	startOctopusLangOrder []string                    //启动OctopusLang
	stopOctopusLangOrder  []string                    //停止OctopusLang
}

//core初始化
func (d *OctopusDirector) directorInit() {
	core.Create()
}

//core启动操作
func (d *OctopusDirector) directorStart() {
	if d.ProfilesCfgPath == nil || len(d.ProfilesCfgPath) == 0 {
		d.ProfilesCfgPath = []string{ProfilesCfgPath}
	}
	if d.ProfilesCfgPrefix == "" {
		d.ProfilesCfgPrefix = ProfilesCfgPrefix
	}
	if d.ProfilesCfgType == "" {
		d.ProfilesCfgType = ProfilesCfgType
	}
	//读配置
	d.Config = loadProfilesCfgPathes(d.ProfilesCfgPath, d.ProfilesCfgPrefix, d.ProfilesCfgType)
	//core加载配置
	core.DependenceInject(d.Config)
	//获得core start注入
	startOctopusLang := core.GetStartOctopusLang()
	d.startOctopusLangOrder = FindDeepOrders(startOctopusLang)
	//调用顺序stop
	//fmt.Println("Start !!!")
	//调用顺序start
	for i := range d.startOctopusLangOrder {
		core.CallStartByName(d.startOctopusLangOrder[i])
	}

	for true {

	}
}

//core停止操作
func (d *OctopusDirector) directorStop() {
	//获得core stop注入
	stopOctopusLang := core.GetStopOctopusLang()
	d.stopOctopusLangOrder = FindHeadOrders(stopOctopusLang)
	//调用顺序stop
	//fmt.Println("Stop !!!")
	for i := range d.stopOctopusLangOrder {
		core.CallStopByName(d.stopOctopusLangOrder[i])
	}

}

func (d *OctopusDirector) setProfilesCfgPath(ProfilesCfgPath string) {
	d.ProfilesCfgPath = append(d.ProfilesCfgPath, ProfilesCfgPath)
}

func (d *OctopusDirector) setProfilesCfgPrefix(ProfilesCfgPrefix string) {
	d.ProfilesCfgPrefix = ProfilesCfgPrefix
}

func (d *OctopusDirector) setProfilesCfgType(ProfilesCfgType string) {
	d.ProfilesCfgType = ProfilesCfgType
}

func (d *OctopusDirector) readCfg(path []string) interface{} {
	Config := d.Config
	for i, p := range path {
		value := Config[p]
		if i+1 == len(path) {
			return value
		} else if strings.Contains(reflect.TypeOf(value).String(), "map") {
			Config = value.(map[interface{}]interface{})
		}
	}
	str := strings.Join(path, ":") + "路劲不对"
	panic(str)
}

func (d *OctopusDirector) writeCfg(writeCfg interface{}, path []string) {
	Config := d.Config
	for i, p := range path {
		value := Config[p]
		if i+1 == len(path) {
			if value != writeCfg {

			}
		} else if strings.Contains(reflect.TypeOf(value).String(), "map") {
			Config = value.(map[interface{}]interface{})
		}
	}
	str := strings.Join(path, ":") + "路劲不对"
	panic(str)
}
