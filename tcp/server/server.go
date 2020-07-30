package main

import (
	"flag"
	"github.com/hslam/socket"
	"github.com/hslam/socket/tcp"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":9999", "-addr=:9999")
	flag.Parse()
}
func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	s := tcp.NewSocket()
	l, err := s.Listen(addr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go func(conn socket.Conn) {
			messages := conn.Messages()
			for {
				msg, err := messages.ReadMessage()
				if err != nil {
					break
				}
				messages.WriteMessage(msg)
			}
			messages.Close()
		}(conn)
	}
}
