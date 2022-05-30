package db

import (
	"octopus/log"
	"octopus/tcp"
)

//TcpClinet回调类
type P2pCallTcpClinet struct {
}

//TcpClinet回调方法
func (l *P2pCallTcpClinet) Call(in interface{}) {
	callJob := in.(tcp.ClinetTcpCallJob)
	msg := callJob.TcpMsg.Msg
	log.Debug(msg)
}

//TcpServer回调类
type P2pCallTcpServer struct {
}

//TcpServer回调方法
func (l *P2pCallTcpServer) Call(in interface{}) {

}
