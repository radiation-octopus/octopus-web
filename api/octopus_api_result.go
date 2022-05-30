package api

//返回的结果：
type Result struct {
	Code int         `json:"code"` //提示代码
	Msg  string      `json:"msg"`  //提示信息
	Data interface{} `json:"data"` //数据
}

//成功
func ResultSuccess(data interface{}) Result {
	res := Result{}
	res.Code = 0
	res.Msg = ""
	res.Data = data
	return res
}

//出错
func ResultError(code int, msg string) Result {
	res := Result{}
	res.Code = code
	res.Msg = msg
	res.Data = ""
	return res
}
