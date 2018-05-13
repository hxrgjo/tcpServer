package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

const (
	ConnectionType = "tcp"
	textQuit       = "quit"
)

type server struct {
	host        string
	port        string
	handlerFunc func(ctx context.Context, c net.Conn) error
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}

		buff, err := ioutil.ReadAll(c)
		if err != nil {
			return
		}

		receiveText := string(buff)

		if receiveText == textQuit {
			cancel()
			c.Close()

			ctx, cancel = context.WithCancel(context.Background())

			continue
		}

		err = s.handlerFunc(ctx, c)
	}
}

func (s *server) HandleFunc(f func(ctx context.Context, c net.Conn) error) {
	s.handlerFunc = f
}

func BasicWebServer(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!!")

		//simulate time out
		//time.Sleep(10 * time.Second)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
