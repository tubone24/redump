package utils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Api struct {
	timeout    int
	httpclient *http.Client
}

func NewHttpClient(timeout int, opts ...Option) *Api {
	api := &Api{
		timeout:    timeout,
		httpclient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(api)
	}

	return api
}

type Option func(*Api)

func OptionHTTPClient(c *http.Client) Option {
	return func(api *Api) {
		api.httpclient = c
	}
}

func (api *Api) request(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)

	respCh := make(chan *http.Response)
	errCh := make(chan error)

	go func() {
		resp, err := api.httpclient.Do(req)
		if err != nil {
			errCh <- err
			return
		}

		respCh <- resp
	}()

	select {
	case resp := <-respCh:
		return resp, nil

	case err := <-errCh:
		return nil, err

	case <-ctx.Done():
		return nil, errors.New("HTTP request cancelled")
	}
}

type HttpClientError struct {
	StatusCode int
}

func (e *HttpClientError) Error() string {
	return "HTTP Client error!: " + strconv.Itoa(e.StatusCode)
}

func (api *Api) Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(api.timeout)*time.Millisecond)
	defer cancel()
	resp, err := api.request(ctx, req)
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

func (api *Api) Post(url, contentType string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(api.timeout)*time.Millisecond)
	defer cancel()
	req.Header.Set("Content-Type", contentType)
	resp, err := api.request(ctx, req)
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

func (api *Api) Put(url, contentType string, data []byte) error {
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(api.timeout)*time.Millisecond)
	defer cancel()
	resp, err := api.request(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		fmt.Println(resp.StatusCode)
		return &HttpClientError{StatusCode: resp.StatusCode}
	}
	return nil
}

func (api *Api) Delete(url string) error {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(api.timeout)*time.Millisecond)
	defer cancel()
	resp, err := api.request(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		fmt.Println(resp.StatusCode)
		return &HttpClientError{StatusCode: resp.StatusCode}
	}
	return nil
}

func NewProxyClient(proxyUrl string) (*http.Client, error) {
	pu, err := url.Parse(proxyUrl)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(pu),
		},
	}
	return client, nil
}
