package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"tcpServer/request"
	"tcpServer/server"
)

var (
	Host = flag.String("h", "localhost", "host")
	Port = flag.String("p", "16680", "port")

	RequestPerSecond int32 = 30
)

func main() {

	//mock api server
	go server.BasicWebServer("8080")

	s := server.New(*Host, *Port)
	s.RequestRateLimit = RequestPerSecond

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
