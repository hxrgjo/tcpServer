package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

const (
	ConnectionType = "tcp"
	textQuit       = "quit"
)

var (
	CurrentRequestPerSecond, ProcessedRequestCount *int32
)

type server struct {
	host                    string
	port                    string
	handlerFunc             func(ctx context.Context, c net.Conn) error
	RequestRateLimit        int32
	currentRequestPerSecond int32
	processedRequestCount   int32
	timer                   *time.Ticker
}

func New(host, port string) *server {
	s := server{host: host, port: port, timer: time.NewTicker(5 * time.Second)}

	CurrentRequestPerSecond = &s.currentRequestPerSecond
	ProcessedRequestCount = &s.processedRequestCount

	go s.resetCurrentRequestPerSecond()
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

		//add request count
		atomic.AddInt32(&s.processedRequestCount, 1)

		//check rate
		if atomic.LoadInt32(&s.currentRequestPerSecond) > s.RequestRateLimit {
			fmt.Println("request rate is over limit")
			c.Close()
			continue
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

		atomic.AddInt32(&s.currentRequestPerSecond, 1)

		go s.handlerFunc(ctx, c)
	}
}

func (s *server) HandleFunc(f func(ctx context.Context, c net.Conn) error) {
	s.handlerFunc = f
}

func (s *server) resetCurrentRequestPerSecond() {
	for {
		select {
		case <-s.timer.C:
			atomic.StoreInt32(&s.currentRequestPerSecond, 0)
		}
	}
}

func BasicWebServer(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!!")

		//simulate time out
		//time.Sleep(10 * time.Second)
	})

	http.HandleFunc("/monitor", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"request_rate":             CurrentRequestPerSecond,
			"current_connection_count": CurrentRequestPerSecond,
			"current_request_rate":     CurrentRequestPerSecond,
			"processed_request_count":  ProcessedRequestCount,
		}

		output := output.Output{Data: data, Status: Status{Code: "0", Message: "Success", Datetime: time.Now()}}
		b, _ := json.Marshal(output)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(b))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
