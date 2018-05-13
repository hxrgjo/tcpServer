package main

import (
	"context"
	"fmt"
	"net"
	"tcpServer/request"
	"tcpServer/server"
)

const (
	Host = "localhost"
	Port = "16680"
)

var requestPerSecond int32 = 30

func main() {

	//mock api server
	go server.BasicWebServer("8080")

	s := server.New(Host, Port)
	s.RequestRateLimit = requestPerSecond

	s.HandleFunc(callbackWithContext)
	s.Listen()
}

func callbackWithContext(ctx context.Context, c net.Conn) (err error) {
	defer c.Close()

	select {
	case <-ctx.Done():
		return
	default:
		url := "http://localhost:8080"
		err := request.SendRequestWithContext(ctx, url)
		if err != nil {
			fmt.Printf("%v\n", err)
			return err
		}
	}

	return
}
