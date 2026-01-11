package handler

import (
	"fmt"
	"io"
	"net"
)

func HandleClient(conn net.Conn) {
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
