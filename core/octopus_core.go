package core

import (
	"reflect"
	"strings"
)

type OctopusCore struct {
	OctopusLang      map[string]*interface{}     //初始化注入结构体
	Config           map[interface{}]interface{} //配置文件详情
	startOctopusLang map[string][]string         //启动OctopusLang
	stopOctopusLang  map[string][]string         //停止OctopusLang
}

//注册
func (c *OctopusCore) register(lang interface{}) {
	name := reflect.TypeOf(lang).String()
	c.OctopusLang[name] = &lang
}

//取消注册
func (c *OctopusCore) unregister(lang interface{}) {
	delete(c.OctopusLang, reflect.TypeOf(&lang).String())
}

//初始化octopusLang
func (c *OctopusCore) create() {
	c.OctopusLang = make(map[string]*interface{})
	c.Config = make(map[interface{}]interface{})
	c.startOctopusLang = make(map[string][]string)
	c.stopOctopusLang = make(map[string][]string)
}

//销毁octopusLang
func (c *OctopusCore) destroy() {
	c.OctopusLang = nil
}

//依赖注入octopusLang
func (c *OctopusCore) dependenceInject() {
	OctopusLang := c.OctopusLang
	for _, v := range OctopusLang {
		valueOf := reflect.ValueOf(*v)
		typeOf := reflect.TypeOf(*v)
		valueOfElem := valueOf.Elem()
		typeOfElem := typeOf.Elem()
		//需要依赖的
		autoRelyonLangs := []string{}
		for i := 0; i < typeOfElem.NumField(); i++ {
			field := valueOfElem.Field(i)
			tag := typeOfElem.Field(i).Tag
			//注入Lang
			if tag.Get(autoInjectLangTag) != "" {
				attribute := OctopusLang["*"+tag.Get(autoInjectLangTag)]
				field.Set(reflect.ValueOf(*attribute))
				//fmt.Println(field)
				//依赖Lang
			} else if tag.Get(autoRelyonLangTag) != "" {
				autoRelyonLangs = append(autoRelyonLangs, "*"+tag.Get(autoRelyonLangTag))
				attribute := OctopusLang["*"+tag.Get(autoRelyonLangTag)]
				field.Set(reflect.ValueOf(*attribute))
				//fmt.Println(field)
				//配置文件yaml
			} else if tag.Get(autoInjectCfgTag) != "" {
				strs := strings.Split(tag.Get(autoInjectCfgTag), ".")
				long := c.readCfg(strs)
				attribute := reflect.ValueOf(long)
				field.Set(attribute)
				//fmt.Println(field)
			}
		}
		startMethod := valueOf.MethodByName(startMethodName)
		stopMethod := valueOf.MethodByName(stopMethodName)
		//启动器
		if startMethod.String() != "<invalid Value>" {
			c.startOctopusLang[typeOf.String()] = autoRelyonLangs
			//startMethod.Call(nil)
			//停止器
		} else if stopMethod.String() != "<invalid Value>" {
			c.stopOctopusLang[typeOf.String()] = autoRelyonLangs
			//stopMethod.Call(nil)
		}
	}
}

//读取OctopusLang
func (c *OctopusCore) getLang(langName string) interface{} {
	key := "*" + langName
	attribute := c.OctopusLang[key]
	return *attribute
}

//写配置
func (c *OctopusCore) setCfg(Config map[interface{}]interface{}) {
	c.Config = Config
}

//查配置
func (c *OctopusCore) readCfg(path []string) interface{} {
	Config := c.Config
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

func (c *OctopusCore) longsByStartMethod() map[string][]string {
	return c.startOctopusLang
}

func (c *OctopusCore) longsByStopMethod() map[string][]string {
	return c.stopOctopusLang
}

func (c *OctopusCore) CallMethod(name string, methodName string, in []interface{}) []interface{} {
	outInterface := []interface{}{}
	octopusLang := c.OctopusLang[name]
	valueOf := reflect.ValueOf(*octopusLang)
	//fmt.Println("valueOf====> ", valueOf)
	method := valueOf.MethodByName(methodName)
	if in == nil || len(in) == 0 || in[0] == nil {
		method.Call(nil)
		//fmt.Println(method)
	} else {
		values := make([]reflect.Value, len(in))
		for i := range in {
			val := in[i]
			value := reflect.ValueOf(val)
			//fmt.Println("value====> ", value)
			values[i] = value
		}

		callValues := method.Call(values)
		for i, c := range callValues {
			outInterface[i] = c.Interface()
		}
		//fmt.Println(method)
	}
	return outInterface
}
