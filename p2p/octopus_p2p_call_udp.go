package db

import (
	"encoding/json"
	"octopus/udp"
	"strconv"
	"time"
)

//udp回调类
type P2pCallUdp struct {
}

//udp回调方法
func (l *P2pCallUdp) Call(in interface{}) {
	job := in.(udp.UdpCallJob)
	jsonStr := job.Msg
	msg := &Msg{}
	json.Unmarshal([]byte(jsonStr), &msg)
	switch msg.TypeName {
	case HeartMsgName:
		heartMsg := (*(msg.MsgInfo)).(HeartMsg)
		heartCall(job.Ip, job.Port, job.Time, heartMsg)
	case DelayMsgName:
		delayMsg := (*(msg.MsgInfo)).(DelayMsg)
		delayCall(job.Ip, job.Port, delayMsg)
	}
}

//心跳数据回调
func heartCall(ip string, udpPort int, heartTime time.Time, msg HeartMsg) {
	host := ip + ":" + strconv.Itoa(udpPort)
	instance := HeartMap[host]
	if instance == nil {
		in := new(Instance)
		in.Ip = ip
		in.UdpPort = udpPort
		in.State = msg.State
		in.HeartTime = heartTime
		HeartMap[host] = in
		updateHeartState(instance)
	} else if instance.State == msg.State {
		instance.HeartTime = heartTime
	} else {
		instance.State = msg.State
		instance.HeartTime = heartTime
		updateHeartState(instance)
	}
}

func updateHeartState(in *Instance) {
	switch in.State {

	case HeartStatePending:

	case HeartStateAvailable:

	case HeartStateAssigned:

	case HeartStateOnline:

	}
}

//延时数据回调
func delayCall(ip string, udpPort int, msg DelayMsg) {

}
