package utils

import (
	"sync"
)

type Pool struct {
	mu       sync.Mutex
	maxJobs  int        // 最大连接数
	jobsChan chan *Jobs //当前池中任务数量
	close    bool
}

func NewPool(max int) *Pool {
	p := &Pool{
		maxJobs:  max,
		jobsChan: make(chan *Jobs, max),
		close:    false,
	}
	return p
}

type Jobs interface {
	Close()
	Execute()
}

func (p *Pool) Start() {
	go func() {
		for true {
			select {
			case job := <-p.jobsChan:
				jobs := *job
				go jobs.Execute()
				//fmt.Println("received", job)
			}
		}
	}()
}

// 将连接返回池中
func (p *Pool) PutJobs(d Jobs) {
	if p.close {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.jobsChan <- &d
}

// 关闭池
func (p *Pool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()
	for d := range p.jobsChan {
		v := *d
		v.Close()
	}
	p.close = true
}
