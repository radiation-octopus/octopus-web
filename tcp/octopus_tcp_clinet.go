package tcp

import (
	"net"
	"octopus/core"
	"octopus/log"
	"strconv"
	"time"
)

//Tcp接受消息回调方法
type ClinetTcpCallJob struct {
	Time   time.Time
	TcpMsg *TcpMsg
}

type TcpMsg struct {
	Ip   string
	Port int
	Msg  string
}

type OctopusClinetTcp struct {
	TcpClinetPort int
	TcpMsgChan    chan *TcpMsg
}

func (t *OctopusClinetTcp) sendGoroutines() {
	go func() {
		for {
			select {
			case tcpMsg := <-t.TcpMsgChan:
				buf := []byte(tcpMsg.Msg)
				host := tcpMsg.Ip + ":" + strconv.Itoa(tcpMsg.Port)
				netAddr := &net.TCPAddr{Port: t.TcpClinetPort}
				dialer := &net.Dialer{LocalAddr: netAddr}
				conn, _ := dialer.Dial("tcp", host)
				conn.Write(buf)
				n, _ := conn.Read(buf[:])
				if n != 0 {
					log.Info(string(buf[:n]))
					//发送消息后接受消息回调
					callJob := new(ClinetTcpCallJob)
					callJob.TcpMsg = tcpMsg
					callJob.Time = time.Now()
					core.CallMethod(TcpClinetAcceptCallBindingStruct, TcpClinetAcceptCallBindingMethod, callJob)
					conn.Close()
				}
			}
		}
	}()
}

//
func (t *OctopusClinetTcp) sendMsg(tcpMsg *TcpMsg) {
	t.TcpMsgChan <- tcpMsg
}

func (t *OctopusClinetTcp) close() {
	close(t.TcpMsgChan)
}

func (t *OctopusClinetTcp) start() {
	t.TcpClinetPort = TcpClinetPort
	t.TcpMsgChan = make(chan *TcpMsg, TcpClinetMsgNum)
	t.sendGoroutines()
}
