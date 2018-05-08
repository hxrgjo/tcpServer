package main

import (
	"fmt"
	"net"
)

const (
	Host           = "localhost"
	Port           = "16680"
	ConnectionType = "tcp"
)

func main() {
	l, err := net.Listen(ConnectionType, Host+":"+Port)
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}

		//deal with handler
		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	fmt.Println("get the tcp request")

	c.Write([]byte("test"))
}
