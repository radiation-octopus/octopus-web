package tcp

import (
	"bufio"
	"fmt"
	"net"
	"octopus/core"
	"octopus/log"
	"octopus/utils"
	"strconv"
	"strings"
	"time"
)

//Tcp接受消息回调方法
type ServerTcpCallJob struct {
	Ip   string
	Port int
	Msg  string
	Time time.Time
	Conn *net.Conn
}

//关闭
func (c *ServerTcpCallJob) Close() {

}

//执行方法
func (c *ServerTcpCallJob) Execute() {
	conn := *(c.Conn)
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	n, _ := reader.Read(buf[:]) // 读取数据
	log.Info(conn.RemoteAddr().String())
	host := strings.Split(conn.RemoteAddr().String(), ":")
	Msg := string(buf[:n])

	c.Msg = Msg
	c.Ip = host[0]
	c.Port, _ = strconv.Atoi(host[1])
	c.Time = time.Now()
	//不用线程池 因为面向链接
	if TcpServerAcceptCallBindingMethod == "" {
		log.Info("CallJob Execute===>> ", c)
	} else {
		core.CallMethod(TcpServerAcceptCallBindingStruct, TcpServerAcceptCallBindingMethod, c)
	}
	conn.Close()
}

type OctopusServerTcp struct {
	TcpServerPort int
	listen        *net.TCPListener
	pool          *utils.Pool
}

func (t *OctopusServerTcp) stop() {
	t.listen.Close()
}

func (t *OctopusServerTcp) start() {
	t.TcpServerPort = TcpServerPort
	//创建pool
	t.pool = utils.NewPool(TcpServerAcceptCallBindingPoolNum)
	t.pool.Start()
	//t.TcpMsgChan = make(chan *TcpMsg, 1024)
	var err error
	t.listen, err = net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: t.TcpServerPort,
	})
	if err != nil {
		fmt.Printf("listen failed,err:%v\n", err)
		return
	}
	t.accpectGoroutines()
	//t.sendGoroutines()
}

func (t *OctopusServerTcp) accpectGoroutines() {
	go func() {
		for {
			conn, _ := t.listen.Accept()
			callJob := new(ServerTcpCallJob)
			callJob.Conn = &conn
			t.pool.PutJobs(callJob)
		}
	}()
}

func (t *OctopusServerTcp) sendGoroutines(msg string, conn net.Conn) {
	buf := []byte(msg)
	conn.Write(buf)
}

func (t *OctopusServerTcp) close() {
	t.pool.Close()
}
