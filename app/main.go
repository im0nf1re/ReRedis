package main

import (
	"fmt"
	"io"
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

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	readBuf := make([]byte, 1024)
	for {
		n, err := conn.Read(readBuf)

		if err == io.EOF {
			fmt.Println("Client disconnected")
			return
		}

		if err != nil {
			fmt.Println("Error reading request: ", err.Error())
			return
		}

		if n == 0 {
			continue
		}

		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			fmt.Println("Error responding: ", err.Error())
			return
		}
	}
}
