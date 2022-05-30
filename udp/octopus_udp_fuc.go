package udp

import (
	"sync"
)

var octopusUdp *OctopusUdp

var once sync.Once

//单例模式
func getInstance() *OctopusUdp {
	once.Do(func() {
		octopusUdp = new(OctopusUdp)
	})
	return octopusUdp
}

func SendMsg(udpMsg *UdpMsg) {
	getInstance().sendMsg(udpMsg)
}

func Start() {
	getInstance().start()
}

func Stop() {
	getInstance().close()
	octopusUdp = nil
}
