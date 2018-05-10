package request

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	GET  = "GET"
	POST = "POST"
)

func Get(url string, headers map[string]string, form url.Values) ([]byte, error) {

	if len(form.Encode()) > 0 {
		url = url + "?" + form.Encode()
	}

	return sendRequest(url, headers, nil, GET)
}

func Post(url string, headers map[string]string, form url.Values) ([]byte, error) {
	return sendRequest(url, headers, form, POST)
}

func sendRequest(url string, headers map[string]string, form url.Values, method string) ([]byte, error) {
	var empty []byte

	client := &http.Client{
		Timeout: time.Duration(600 * time.Second),
	}

	var formBody io.Reader
	if form != nil {
		formBody = strings.NewReader(form.Encode())
	}

	req, err := http.NewRequest(method, url, formBody)
	if err != nil {
		return empty, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return empty, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return empty, err
	}

	return respBody, nil
}
