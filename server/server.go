package server

import (
	"fmt"
	"net"
)

const ConnectionType = "tcp"

type server struct {
	host        string
	port        string
	handlerFunc func(c net.Conn)
}

type Handler struct {
}

func New(host, port string) *server {
	s := server{host: host, port: port}
	return &s
}

func (s *server) Listen() {

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
		s.handlerFunc(c)
	}
}

func (s *server) HandleFunc(f func(net.Conn)) {
	s.handlerFunc = f
}
