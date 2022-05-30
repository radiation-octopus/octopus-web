package udp

import (
	"fmt"
	"net"
	"octopus/core"
	"octopus/log"
	"octopus/utils"
	"strconv"
	"strings"
	"time"
)

//Udp接受消息回调方法
type UdpCallJob struct {
	Ip   string
	Port int
	Msg  string
	Time time.Time
}

//关闭
func (c *UdpCallJob) Close() {

}

//执行方法
func (c *UdpCallJob) Execute() {
	if UdpAcceptCallBindingMethod == "" {
		log.Info("CallJob Execute===>> ", c)
	} else {
		core.CallMethod(UdpAcceptCallBindingStruct, UdpAcceptCallBindingMethod, c)
	}
}

type UdpMsg struct {
	Ip   string
	Port int
	Msg  string
}

type OctopusUdp struct {
	Port       int
	listen     *net.UDPConn
	pool       *utils.Pool
	UdpMsgChan chan *UdpMsg
}

func (u *OctopusUdp) stop() {
	u.listen.Close()
}

func (u *OctopusUdp) start() {
	u.Port = UdpPort
	//创建pool
	u.pool = utils.NewPool(UdpAcceptCallBindingPoolNum)
	u.pool.Start()
	u.UdpMsgChan = make(chan *UdpMsg, UdpMsgNum)
	var err error
	u.listen, err = net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: u.Port,
	})
	if err != nil {
		fmt.Printf("listen failed,err:%v\n", err)
		return
	}
	u.accpectGoroutines()
	u.sendGoroutines()
}

func (u *OctopusUdp) accpectGoroutines() {
	go func() {
		for {
			var buf [1024]byte
			n, addr, err := u.listen.ReadFromUDP(buf[:])
			if err != nil {
				//fmt.Printf("read udp failed,err:%v\n", err)
				return
			}
			Msg := string(buf[:n])
			callJob := new(UdpCallJob)
			callJob.Ip = addr.IP.String()
			callJob.Port = addr.Port
			callJob.Msg = Msg
			callJob.Time = time.Now()
			u.pool.PutJobs(callJob)
			//fmt.Printf("接收到的数据:%v\n", " addr:", addr, " str:", Msg)
		}
	}()
}

func (u *OctopusUdp) sendGoroutines() {
	go func() {
		for {
			select {
			case udpMsg := <-u.UdpMsgChan:
				buf := []byte(udpMsg.Msg)
				Ips := strings.Split(udpMsg.Ip, ".")
				a, _ := strconv.Atoi(Ips[0])
				b, _ := strconv.Atoi(Ips[1])
				c, _ := strconv.Atoi(Ips[2])
				d, _ := strconv.Atoi(Ips[3])
				addr := &net.UDPAddr{
					IP:   net.IPv4(byte(a), byte(b), byte(c), byte(d)),
					Port: udpMsg.Port,
				}
				_, err := u.listen.WriteToUDP(buf, addr)
				if err != nil {
					//fmt.Printf("write to %v failed,err:%v\n", addr, err)
					return
				}
			}
		}
	}()
}

func (u *OctopusUdp) sendMsg(udpMsg *UdpMsg) {
	u.UdpMsgChan <- udpMsg
}

func (u *OctopusUdp) close() {
	u.pool.Close()
	close(u.UdpMsgChan)
}
