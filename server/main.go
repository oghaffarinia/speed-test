package main

import (
	"fmt"
	"github.com/omidplus/arvan/config"
	"github.com/omidplus/arvan/protocol"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime)
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <path-to-config-file>", os.Args[0])
	}
	config, err := config.LoadServer(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.Listen(config.Protocol, fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(conn net.Conn) {
			log.Println(conn.RemoteAddr())
			conn.SetDeadline(time.Now().Add(config.Timeout))
			protocol.Handle(conn)
		}(conn)
	}
}
