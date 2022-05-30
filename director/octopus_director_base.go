package director

//启动包需要start继承该接口
type BaseStart interface {
	Start()
}

//启动包需要stop继承该接口
type BaseStop interface {
	Stop()
}
