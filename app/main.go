package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	readBuf := make([]byte, 0)
	for {
		_, err = conn.Read(readBuf)
		if err != nil {
			fmt.Println("Error reading request: ", err.Error())
			os.Exit(1)
		}
		if err == io.EOF {
			break
		}

		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			fmt.Println("Error responding: ", err.Error())
			os.Exit(1)
		}
	}

}
