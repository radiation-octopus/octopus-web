package director

//初始化octopus
func init() {
	//创建核心
	getInstance().directorInit()
	//把启动注入
	Register(new(DirectorStart))
	//把停止注入
	Register(new(DirectorStop))
}
