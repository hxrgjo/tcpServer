package request

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	GET = "GET"
)

func SendRequestWithContext(ctx context.Context, url string) (err error) {

	endSingal := make(chan interface{}, 1)

	tr := &http.Transport{}
	client := &http.Client{
		Timeout:   time.Duration(5 * time.Second),
		Transport: tr,
	}

	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		return
	}

	go func() {
		resp, err := client.Do(req)
		endSingal <- err
		if err != nil {
			return
		}
		respBody, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("get resp.body: %s \n", string(respBody))
		defer resp.Body.Close()
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		fmt.Printf("ctx cancel, err: %v\n", ctx.Err())
		return ctx.Err()
	case err := <-endSingal:
		fmt.Printf("get end singal, err: %v\n", err)
	}

	return
}
