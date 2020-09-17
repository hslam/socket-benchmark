package main

import (
	"flag"
	"github.com/hslam/socket"
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
	go http.ListenAndServe(":6060", nil)
	sock := socket.NewTCPSocket(nil)
	l, err := sock.Listen(addr)
	if err != nil {
		log.Fatalln(err)
	}
	l.ServeMessages(func(messages socket.Messages) (socket.Context, error) {
		return messages, nil
	}, func(context socket.Context) error {
		messages := context.(socket.Messages)
		msg, err := messages.ReadMessage()
		if err != nil {
			return err
		}
		return messages.WriteMessage(msg)
	})
}
