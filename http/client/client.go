package main

import (
	"flag"
	"fmt"
	"github.com/hslam/socket"
	"github.com/hslam/stats"
	"log"
)

var addr string
var clients int
var total int
var msg int
var bar bool

func init() {
	flag.StringVar(&addr, "addr", ":9999", "-addr=:9999")
	flag.IntVar(&total, "total", 100000, "-total=100000")
	flag.IntVar(&clients, "clients", 1, "-clients=1")
	flag.IntVar(&msg, "msg", 512, "-msg=512")
	flag.BoolVar(&bar, "bar", true, "-bar=true")
	flag.Parse()
	stats.SetBar(bar)
	fmt.Printf("./client -addr=%s -total=%d -clients=%d\n", addr, total, clients)
}

func main() {
	if clients < 1 || total < 1 {
		return
	}
	var wrkClients []stats.Client
	var msg = make([]byte, msg)
	for i := 0; i < clients; i++ {
		sock := socket.NewHTTPSocket(nil)
		if conn, err := sock.Dial(addr); err != nil {
			log.Fatalln("dailing error: ", err)
		} else {
			wrkClients = append(wrkClients, &WrkClient{conn.Messages(), msg})
		}
	}
	stats.StartPrint(1, total, wrkClients)
}

type WrkClient struct {
	socket.Messages
	msg []byte
}

func (c *WrkClient) Call() (int64, int64, bool) {
	err := c.Messages.WriteMessage(c.msg)
	if err != nil {
		return 0, 0, false
	}
	data, err := c.Messages.ReadMessage()
	if err != nil {
		return int64(len(c.msg)), 0, false
	}
	return int64(len(c.msg)), int64(len(data)), true
}
