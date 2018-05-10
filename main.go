package main

import (
	"fmt"
	"io/ioutil"
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

	buff, _ := ioutil.ReadAll(c)
	s := string(buff[:])

	fmt.Println(s)

	c.Write([]byte("test"))
}
