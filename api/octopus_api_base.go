package api

//接口api继承
type BaseApi interface {
	//绑定方法
	Binding()
}

type BaseApiObserver interface {
	//添加Api
	AddApi(baseApis ...BaseApi)
	//绑定api method方法
	BindingApI(baseApi BaseApi, f func(map[string]interface{}) map[string]interface{}, path ...string)
}
