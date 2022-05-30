package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"octopus/db"
	"octopus/director"
	"octopus/tcp"
	"octopus/udp"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestUdp(t *testing.T) {
	udpStart := udp.UdpStart{Port: 40000}
	udpStart.Start()
	udpMsg := udp.UdpMsg{Ip: "127.0.0.1", Port: 30000, Msg: "mc_cao"}
	udp.SendMsg(&udpMsg)
	time.Sleep(1 * time.Second)
}

func TestTcp(t *testing.T) {
	tcpStart := tcp.TcpStart{ServerPort: 40000}
	tcpStart.Start()
	tcpMsg := tcp.TcpMsg{Ip: "127.0.0.1", Port: 20000, Msg: "mc_cao"}
	tcp.SendMsg(&tcpMsg)
	time.Sleep(1 * time.Second)
}

//生成token值
func TestToken(t *testing.T) {
	length := 256
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			rs = append(rs, string(rune(rand.Intn(26)+65)))
		} else {
			rs = append(rs, string(rune(rand.Intn(26)+97)))
		}
	}
	fmt.Println(strings.Join(rs, ""))
}

func TestUuid(t *testing.T) {
	// 创建
	u1 := uuid.NewV4()
	fmt.Println(u1.String())
	// 解析
	u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	if err != nil {
		fmt.Println("Something gone wrong: ", err)
		return
	}
	fmt.Println(u2.String())
}

func TestDb(t *testing.T) {
	dbStart := db.DbStart{SavePath: "data/octopus.db"}
	dbStart.Start()
	type User struct {
		ID   int
		Name string
		Age  int
		Sex  string
	}

	u := User{}
	u.ID = 1
	u.Name = "jeffcail"
	u.Age = 18
	u.Sex = "男"

	db.Insert("0201", "0001", &u)
	user2 := db.Query("0201", "0001", &u).(*User)
	fmt.Println(user2)
}

func TestTree(t *testing.T) {
	mapVar := make(map[string][]string)
	mapVar["b"] = []string{"a"}
	mapVar["c"] = []string{"a"}
	mapVar["d"] = []string{"c"}
	fmt.Println(director.FindDeepOrders(mapVar))
	fmt.Println(director.FindHeadOrders(mapVar))
}
