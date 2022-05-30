package tcp

//存放默认常量

//Server 端口
var TcpServerPort int

//Server 回调线程数量
var TcpServerAcceptCallBindingPoolNum int

//Server tcp接受回调方法
var TcpServerAcceptCallBindingMethod string

//Server tcp接受回调注入体
var TcpServerAcceptCallBindingStruct string

//Clinet 端口
var TcpClinetPort int

//Clinet 发送消息队列
var TcpClinetMsgNum int

//Clinet tcp接受回调方法
var TcpClinetAcceptCallBindingMethod string

//Clinet tcp接受回调注入体
var TcpClinetAcceptCallBindingStruct string
