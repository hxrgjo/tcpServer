package request

import (
	"context"
	"tcpServer/server"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestSendRequestWithContext(t *testing.T) {
	go server.BasicWebServer("8080")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	url := "http://localhost:8080"
	go SendRequestWithContext(ctx, url)

	time.Sleep(1 * time.Second)

	assert.Equal(t, 1, 2)
}
