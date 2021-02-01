package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpClientError struct {
	StatusCode int
}

func (e *HttpClientError) Error() string {
	return "HTTP Client error!"
}

func Get(url string) ([]byte, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, &HttpClientError{StatusCode: resp.StatusCode}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func Post(url, contentType string, data []byte) ([]byte, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		fmt.Println(resp.StatusCode)
		return nil, &HttpClientError{StatusCode: resp.StatusCode}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func Put(url, contentType string, data []byte) error {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		fmt.Println(resp.StatusCode)
		return &HttpClientError{StatusCode: resp.StatusCode}
	}
	return nil
}
