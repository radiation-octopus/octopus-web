package db

import (
	"time"
)

//db 数据
type Msg struct {
	TypeName string
	MsgInfo  *interface{}
	SendUuid string
}

//db udp 心跳数据
type HeartMsg struct {
	State string
}

//db udp 延时数据
type DelayMsg struct {
	SendDelayTime time.Time
	SendDelayUuid string
}

//db tcp 详情数据
type DetailsMsg struct {
}

//db 实例
type Instance struct {
	Uuid          string
	Ip            string
	UdpPort       int
	TcpClinetPort int
	TcpServerPort int
	State         string
	HeartTime     time.Time
}
