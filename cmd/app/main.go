package main

import (
	"ReRedis/internal/handler"
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error handling connection:", err)
			continue
		}

		go handler.HandleClient(conn)
	}
}
