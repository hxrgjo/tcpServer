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
	s.HandleFunc(callback)
	s.Listen()
}

func callback(c net.Conn) {
	defer c.Close()

	buff, _ := ioutil.ReadAll(c)
	s := string(buff[:])

	fmt.Println(s)
}
