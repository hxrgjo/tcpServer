package main

import (
	"fmt"
	"net"
	"tcpServer/server"
)

const (
	Host = "localhost"
	Port = "16680"
)

func main() {
	s := server.New(Host, Port)
	s.Listen(callback)
}

func callback(c net.Conn) {
	defer c.Close()

	fmt.Println("get the tcp request")

	c.Write([]byte("test"))
}
