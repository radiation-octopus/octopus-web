package tcp

import "net"

type OctopusTcp struct {
	OctopusClinetTcp *OctopusClinetTcp
	OctopusServerTcp *OctopusServerTcp
}

func (t *OctopusTcp) start() {
	t.OctopusClinetTcp = new(OctopusClinetTcp)
	t.OctopusClinetTcp.start()
	t.OctopusServerTcp = new(OctopusServerTcp)
	t.OctopusServerTcp.start()
}

func (t *OctopusTcp) sendMsg(tcpMsg *TcpMsg) {
	t.OctopusClinetTcp.sendMsg(tcpMsg)
}

func (t *OctopusTcp) close() {
	t.OctopusClinetTcp.close()
	t.OctopusServerTcp.close()
}

func (t *OctopusTcp) sendMsgByConn(msg string, conn net.Conn) {
	t.OctopusServerTcp.sendGoroutines(msg, conn)
}
