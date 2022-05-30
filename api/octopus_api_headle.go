package api

import (
	"net/http"
	"octopus/utils"
)

type OctopusApiHandle struct {
	pool *utils.Pool
}

func initHandle() *OctopusApiHandle {
	h := new(OctopusApiHandle)
	h.pool = utils.NewPool(ApiHandlePoolNum)
	h.pool.Start()
	http.HandleFunc("/", h.httpHandle)
	return h
}

//http监听回调方法
func (h *OctopusApiHandle) httpHandle(writer http.ResponseWriter, request *http.Request) {
	callJob := new(OctopusApiCallJob)
	callJob.writer = writer
	callJob.request = request
	h.pool.PutJobs(callJob)
}

type OctopusApiCallJob struct {
	writer  http.ResponseWriter
	request *http.Request
}

func (c *OctopusApiCallJob) Close() {

}

//执行方法
func (c *OctopusApiCallJob) Execute() {
	process := initProcess(c.writer, c.request)
	process.executeProcess()
}
