package tcp

import (
	"net"
	"sync"
)

var octopusTcp *OctopusTcp

var once sync.Once

//单例模式
func getInstance() *OctopusTcp {
	once.Do(func() {
		octopusTcp = new(OctopusTcp)
	})
	return octopusTcp
}

func Start() {
	getInstance().start()
}

func SendMsg(tcpMsg *TcpMsg) {
	getInstance().sendMsg(tcpMsg)
}

func SendMsgByConn(msg string, conn net.Conn) {
	getInstance().sendMsgByConn(msg, conn)
}

func Stop() {
	getInstance().close()
	octopusTcp = nil
}

//
//func TcpAcceptCallBinding(method string){
//	TcpAcceptCallBindingMethod=method
//}
