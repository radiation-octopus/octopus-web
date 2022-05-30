package db

import (
	"time"
)

type P2pMsg struct {
	P2pLevel string
	Time     time.Time
	Contents string
	FuncInfo string
}

type OctopusP2p struct {
	//P2pMsgChan chan *P2pMsg
}

func (p *OctopusP2p) start() {
}

func (p *OctopusP2p) stop() {

}
