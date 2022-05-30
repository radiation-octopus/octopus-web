package blockchain

import (
	"reflect"
	"sync"
)

//Feed实现了一对多订阅，其中事件的载体是一个频道
type Feed struct {
	once      sync.Once        //只能初始化一次
	sendlock  chan struct{}    //单元素缓冲区
	removeSub chan interface{} //中断发送
	sendCases caseList         //元素活动集
	mu        sync.Mutex
	inbox     caseList
	etype     reflect.Type
}

type caseList []reflect.SelectCase
