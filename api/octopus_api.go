package api

import (
	"octopus/director"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

type OctopusApi struct {
	RouteMap          map[string]interface{} //路径map
	StructNameMutex   map[string]*sync.Mutex //结构体锁
	ApiHandle         *OctopusApiHandle      //api回调
	Sessioncollection []*OctopusApiSession   //session集合
}

//启动start
func (a *OctopusApi) start() {
	//core.Create()
	a.RouteMap = make(map[string]interface{})
	a.StructNameMutex = make(map[string]*sync.Mutex)
	if IsSession {
		a.Sessioncollection = []*OctopusApiSession{}
	}
	//初始化ApiHandle
	a.ApiHandle = initHandle()
}

//启动stop
func (a *OctopusApi) stop() {
}

func (a *OctopusApi) AddApi(baseApis ...BaseApi) {
	for _, api := range baseApis {
		a.StructNameMutex[reflect.TypeOf(&api).String()] = &sync.Mutex{}
		//绑定路由
		api.Binding()
		//注入到启动器中
		director.Register(a)
	}
}

func (a *OctopusApi) getMutex(baseApi BaseApi) *sync.Mutex {
	return a.StructNameMutex[reflect.TypeOf(&baseApi).String()]
}

func (a *OctopusApi) BindingApI(
	baseApi BaseApi,
	f func(map[string]interface{}) map[string]interface{},
	requestMethod string,
	path ...string) {
	//获取binding参数
	binding := initApiBinding(baseApi, f)
	typeOfElem := reflect.TypeOf(&baseApi).Elem()
	var baseUrl string
	for i := 0; i < typeOfElem.NumField(); i++ {
		tag := typeOfElem.Field(i).Tag
		if tag.Get(BaseUrlTag) != "" {
			baseUrl = tag.Get(BaseUrlTag)
		}
	}
	basepath := strings.Split(baseUrl, "/")
	var allpath []string
	allpath = basepath
	for _, p := range path {
		allpath = append(allpath, p)
	}
	a.setBinding(allpath, requestMethod, binding)
}

func (a *OctopusApi) setBinding(path []string, requestMethod string, binding *OctopusApiBinding) {
	routeMap := a.RouteMap
	for _, p := range path {
		num := 0
		isHasPathVariable := false
		for k, _ := range routeMap {
			//是否包含方法前缀
			if !strings.Contains(k, RouteMapMethodPrefix) {
				num++
			}
			//是否包含路径值前缀
			if strings.Contains(k, RouteMapPathVariablePrefix) {
				isHasPathVariable = true
			}
		}
		value := routeMap[p]
		//是不是空
		if num == 0 {
			value = make(map[string]interface{})
			routeMap[p] = value
			routeMap = value.(map[string]interface{})
		} else {
			//是不是有路径值
			if isHasPathVariable {
				if value == p {
					routeMap = value.(map[string]interface{})
				} else {
					panic("路径重复，出现异常")
				}
			} else {
				if strings.Contains(p, RouteMapPathVariablePrefix) {
					panic("路径重复，出现异常")
				} else {
					if value == nil {
						value = make(map[string]interface{})
						routeMap[p] = value
						routeMap = value.(map[string]interface{})
					} else {
						routeMap = value.(map[string]interface{})
					}
				}
			}
		}
	}
	routeMap[RouteMapMethodPrefix+requestMethod] = binding
}

func (a *OctopusApi) getBindingAndParam(path []string, requestMethod string, param map[string]interface{}) (*OctopusApiBinding, map[string]interface{}) {
	//param := map[string]interface{}{}
	routeMap := a.RouteMap
	var rkstr string
	for i, p := range path {
		isHasPathVariable := false
		if len(routeMap) == 1 {
			for rk, _ := range routeMap {
				rkstr = rk
				isHasPathVariable = strings.Contains(rk, RouteMapPathVariablePrefix)
			}
		}
		if isHasPathVariable {
			k := strings.Replace(rkstr, RouteMapPathVariablePrefix, "", 1)
			v := p
			fl, errofl := strconv.ParseFloat(v, 64)
			in, erroin := strconv.ParseInt(v, 10, 64)
			if errofl == nil {
				param[k] = fl
			} else if erroin == nil {
				param[k] = in
			} else {
				param[k] = p
			}
			param[k] = v
			value := routeMap[rkstr]
			routeMap = value.(map[string]interface{})
		} else {
			value := routeMap[p]
			if i+1 == len(path) {
				routeMap = value.(map[string]interface{})
				value = routeMap[requestMethod]
				return value.(*OctopusApiBinding), param
			} else if strings.Contains(reflect.TypeOf(value).String(), "map") {
				routeMap = value.(map[string]interface{})
			}
		}

	}
	str := strings.Join(path, ":") + "路劲不对"
	panic(str)
}

//获取session
func (a *OctopusApi) getSession(key string) *OctopusApiSession {
	for _, s := range a.Sessioncollection {
		if s.Key == key {
			return s
		}
	}
	return nil
}

//创建session
func (a *OctopusApi) createSession(key string) {
	session := initSession(key)
	a.Sessioncollection = append(a.Sessioncollection, session)
}

//清理session
func (a *OctopusApi) cleanSession() {
	timeNow := time.Now()
	var session []*OctopusApiSession
	for _, s := range a.Sessioncollection {
		if s.ExpirateTime.After(timeNow) {
			session = append(session, s)
		}
	}
}
