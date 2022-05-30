package api

import (
	"encoding/json"
	"net/http"
	"octopus/utils"
	"reflect"
	"strconv"
	"strings"
)

//执行流程
type OctopusApiProcess struct {
	Writer    http.ResponseWriter    //respone 写入流
	Request   *http.Request          //request 对象
	GetValues []*reflect.Value       //读取参数反射valueof数组
	GetTypes  []*reflect.StructField //读取参数反射类型数组
	SetValues []*reflect.Value       //写入参数反射valueof数组
	SetTypes  []*reflect.StructField //写入参数反射类型数组
	OutMap    map[string]interface{} //回调函数出参
	InMap     map[string]interface{} //回调函数入参
	Binding   *OctopusApiBinding     //Binding 绑定对象
}

//初始化过程
func initProcess(writer http.ResponseWriter, request *http.Request) *OctopusApiProcess {
	apiProcess := new(OctopusApiProcess)
	apiProcess.InMap = make(map[string]interface{})
	apiProcess.OutMap = make(map[string]interface{})
	apiProcess.Writer = writer
	apiProcess.Request = request
	requestURI := request.RequestURI
	uri := strings.Split(requestURI, "?")
	path := strings.Split(uri[0], "/")
	apiProcess.Binding, apiProcess.InMap = getInstance().getBindingAndParam(path, request.Method, apiProcess.InMap)
	if len(uri) == 2 {
		param := strings.Split(uri[1], "&")
		for _, p := range param {
			s := strings.Split(p, "=")
			f, errof := strconv.ParseFloat(s[1], 64)
			i, erroi := strconv.ParseInt(s[1], 10, 64)
			if errof == nil {
				apiProcess.InMap[s[0]] = f
			} else if erroi == nil {
				apiProcess.InMap[s[0]] = i
			} else {
				apiProcess.InMap[s[0]] = s[1]
			}
		}
	}
	//apiProcess.BaseApi
	valueOfElem := reflect.ValueOf(apiProcess.Binding.BaseApi).Elem()
	typeOfElem := reflect.TypeOf(apiProcess.Binding.BaseApi).Elem()
	for i := 0; i < valueOfElem.NumField(); i++ {
		field := valueOfElem.Field(i)
		typeOf := typeOfElem.Field(i)
		tag := typeOf.Tag
		if tag.Get(GetRequestHeadTag) != "" {
			apiProcess.GetValues = append(apiProcess.GetValues, &field)
			apiProcess.GetTypes = append(apiProcess.GetTypes, &typeOf)
		} else if IsCookie && tag.Get(GetCookieTag) != "" {
			apiProcess.GetValues = append(apiProcess.GetValues, &field)
			apiProcess.GetTypes = append(apiProcess.GetTypes, &typeOf)
		} else if tag.Get(GetRequestBodyTag) != "" {
			apiProcess.GetValues = append(apiProcess.GetValues, &field)
			apiProcess.GetTypes = append(apiProcess.GetTypes, &typeOf)
		} else if IsSession && tag.Get(GetSessionTag) != "" {
			apiProcess.GetValues = append(apiProcess.GetValues, &field)
			apiProcess.GetTypes = append(apiProcess.GetTypes, &typeOf)
		} else if tag.Get(SetReponseHeadTag) != "" {
			apiProcess.SetValues = append(apiProcess.SetValues, &field)
			apiProcess.SetTypes = append(apiProcess.SetTypes, &typeOf)
		} else if IsCookie && tag.Get(SetCookieTag) != "" {
			apiProcess.SetValues = append(apiProcess.SetValues, &field)
			apiProcess.SetTypes = append(apiProcess.SetTypes, &typeOf)
		} else if tag.Get(SetReponseBodyTag) != "" {
			apiProcess.SetValues = append(apiProcess.SetValues, &field)
			apiProcess.SetTypes = append(apiProcess.SetTypes, &typeOf)
		} else if IsSession && tag.Get(SetSessionTag) != "" {
			apiProcess.SetValues = append(apiProcess.SetValues, &field)
			apiProcess.SetTypes = append(apiProcess.SetTypes, &typeOf)
		}
	}

	return apiProcess
}

//获值
func (a *OctopusApiProcess) getParam() {
	for i, field := range a.GetValues {
		tag := a.GetTypes[i].Tag
		//获取request 请求头
		if tag.Get(GetRequestHeadTag) != "" {
			headKey := tag.Get(GetRequestHeadTag)
			str := a.Request.Header.Get(headKey)
			attribute := reflect.ValueOf(str)
			field.Set(attribute)
			//获取cookie
		} else if IsCookie && tag.Get(GetCookieTag) != "" {
			headKey := tag.Get(GetRequestHeadTag)
			coolie, _ := a.Request.Cookie(headKey)
			value := coolie.Value
			attribute := reflect.ValueOf(value)
			field.Set(attribute)
			//获取body
		} else if tag.Get(GetRequestBodyTag) != "" {
			bodyKey := tag.Get(GetRequestBodyTag)
			newMap := map[string]interface{}{}
			json.NewDecoder(a.Request.Body).Decode(&newMap)
			//value := coolie.Value
			attribute := reflect.ValueOf(newMap[bodyKey])
			field.Set(attribute)
			//获取session
		} else if IsSession && tag.Get(GetSessionTag) != "" {
			headKey := tag.Get(GetSessionTag)
			coolie, _ := a.Request.Cookie(CookieSessionKeyName)
			if coolie != nil {
				session := getInstance().getSession(coolie.Value)
				if session == nil {
					//CookieSessionValue := utils.GetRandomStr(16)
					getInstance().createSession(coolie.Value)
				} else {
					//m := session.(map[string]interface{})
					//value := m[headKey]
					value := session.getSessionVal(headKey)
					attribute := reflect.ValueOf(value)
					field.Set(attribute)
				}
			} else {
				CookieSessionValue := utils.GetRandomStr(16)
				coolie = &http.Cookie{
					Name:   CookieSessionKeyName,
					Value:  CookieSessionValue,
					Path:   SessionCookiePath,
					Domain: SessionCookieDomain,
					MaxAge: SessionMaxAge,
				}
				getInstance().createSession(coolie.Value)
				http.SetCookie(a.Writer, coolie)
			}
		}
		field.Close()
	}
	//form表单
	a.InMap[FormMapKey] = a.Request.PostForm
}

//写值
func (a *OctopusApiProcess) setParam() {
	for i, field := range a.SetValues {
		tag := a.SetTypes[i].Tag
		//写入 reponse head
		if tag.Get(SetReponseHeadTag) != "" {
			key := tag.Get(SetReponseHeadTag)
			value := utils.GetInToStr(field.Interface())
			a.Writer.Header().Set(key, value)
			//写入 cookie 请求
		} else if IsCookie && tag.Get(SetCookieTag) != "" {
			key := tag.Get(SetCookieTag)
			value := utils.GetInToStr(field.Interface())
			coolie := &http.Cookie{
				Name:   key,
				Value:  utils.GetInToStr(value),
				Path:   CookiePath,
				Domain: CookieDomain,
				MaxAge: CookieMaxAge,
			}
			http.SetCookie(a.Writer, coolie)
			//写入body
		} else if tag.Get(SetReponseBodyTag) != "" {
			key := tag.Get(SetReponseBodyTag)
			value := utils.GetInToStr(field.Interface())
			a.OutMap[key] = value
			//写入session
		} else if IsSession && tag.Get(SetSessionTag) != "" {
			headKey := tag.Get(SetSessionTag)
			coolie, _ := a.Request.Cookie(CookieSessionKeyName)
			if coolie != nil {
				getInstance().getSession(coolie.Value).setSessionVal(headKey, field.Interface())
			}
		}
		//返回值
		field.Close()
	}
	//响应码
	a.Writer.WriteHeader(200)
	j, _ := json.Marshal(ResultSuccess(a.OutMap))
	a.Writer.Write(j)
}

//清理值
func (a *OctopusApiProcess) cleanParam() {
	for _, field := range a.GetValues {
		var str = ""
		field.Set(reflect.ValueOf(str))
		field.Close()
	}
	for _, field := range a.SetValues {
		var str = ""
		field.Set(reflect.ValueOf(str))
		field.Close()
	}
}

//执行回调方法
func (a *OctopusApiProcess) callMethod() {
	callFunc := a.Binding.CallFunc
	a.OutMap = callFunc(a.InMap)
}

//执行流程
func (a *OctopusApiProcess) executeProcess() {
	a.Binding.Mutex.Lock()
	a.getParam()
	a.callMethod()
	a.cleanParam()
	a.setParam()
	a.cleanParam()
	a.Binding.Mutex.Unlock()
}
