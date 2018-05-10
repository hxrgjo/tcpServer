package server

import (
	"fmt"
	"net"
)

const ConnectionType = "tcp"

type server struct {
	host string
	port string
}

func New(host, port string) *server {
	s := server{host: host, port: port}
	return &s
}

func (s *server) Listen(f func(net.Conn)) {

	l, err := net.Listen(ConnectionType, s.host+":"+s.port)
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
		go f(c)
	}
}
