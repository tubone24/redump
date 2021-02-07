package utils_test

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/utils"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

type testResponseJson struct {
	Status string
}

func client(t *testing.T, respTime time.Duration, resp *http.Response) *http.Client {
	t.Helper()

	body := testResponseJson{"OK"}

	b, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	return NewTestClient(func(req *http.Request) *http.Response {
		time.Sleep(respTime)
		if resp != nil {
			return resp
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(b)),
			Header:     make(http.Header),
		}
	})
}

func TestApiGet(t *testing.T) {
	cases := map[string]struct {
		client               *http.Client
		expectHasError       bool
		timeout              int
		expectedErrorMessage string
		expectedText         string
	}{
		"normal": {
			client:         client(t, 0, nil),
			expectHasError: false,
			timeout:        10000,
			expectedText:   "{\"Status\":\"OK\"}",
		},
		"204NoContent": {
			client: client(t, 0, &http.Response{
				StatusCode: http.StatusNoContent,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("{\"Status\":\"OK\"}"))),
				Header:     make(http.Header),
			}),
			expectHasError: false,
			timeout:        10000,
			expectedText:   "{\"Status\":\"OK\"}",
		},
		"timeout": {
			client:               client(t, 0, nil),
			expectHasError:       true,
			timeout:              0,
			expectedErrorMessage: "HTTP request cancelled",
		},
		"500Error": {
			client: client(t, 0, &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("error"))),
				Header:     make(http.Header),
			}),
			expectHasError:       true,
			timeout:              10000,
			expectedErrorMessage: "HTTP Client error!: 500",
		},
		"400Error": {
			client: client(t, 0, &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("error"))),
				Header:     make(http.Header),
			}),
			expectHasError:       true,
			timeout:              10000,
			expectedErrorMessage: "HTTP Client error!: 400",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			e := utils.NewHttpClient(c.timeout, utils.OptionHTTPClient(c.client))

			resp, err := e.Get("https://example.com")

			if c.expectHasError {
				if err == nil {
					t.Errorf("expected error but no errors ouccured")
					return
				}

				if err.Error() != c.expectedErrorMessage {
					t.Errorf("unexpected error message. expected '%s', actual '%s'", c.expectedErrorMessage, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf(err.Error())
				return
			}

			if string(resp) != c.expectedText {
				t.Errorf("unexpected response's text. expected '%s', actual '%s'", c.expectedText, string(resp))
			}
		})
	}
}

func TestApiGetInvalidUrl(t *testing.T) {
	e := utils.NewHttpClient(1000)
	_, err := e.Get("https://jdooudnpoowudqw.jdpofufoqf9-sjhfq-wdwhqwajasjpdjwdjqwoddjqwopdjdqdudl.com.kjpwf.coaa.coomwef.com")
	if err == nil {
		t.Errorf("expected error but no errors ouccured")
	}
}

func TestApiPostInvalidUrl(t *testing.T) {
	e := utils.NewHttpClient(1000)
	_, err := e.Post("https://jdooudnpoowudqw.jdpofufoqf9-sjhfq-wdwhqwajasjpdjwdjqwoddjqwopdjdqdudl.com.kjpwf.coaa.coomwef.com", "application/json", []byte("{\"aaa\": \"aaa\"}"))
	if err == nil {
		t.Errorf("expected error but no errors ouccured")
	}
}

func TestApiPutInvalidUrl(t *testing.T) {
	e := utils.NewHttpClient(1000)
	err := e.Put("https://jdooudnpoowudqw.jdpofufoqf9-sjhfq-wdwhqwajasjpdjwdjqwoddjqwopdjdqdudl.com.kjpwf.coaa.coomwef.com", "application/json", []byte("{\"aaa\": \"aaa\"}"))
	if err == nil {
		t.Errorf("expected error but no errors ouccured")
	}
}

func TestApiDeleteInvalidUrl(t *testing.T) {
	e := utils.NewHttpClient(1000)
	err := e.Delete("https://jdooudnpoowudqw.jdpofufoqf9-sjhfq-wdwhqwajasjpdjwdjqwoddjqwopdjdqdudl.com.kjpwf.coaa.coomwef.com")
	if err == nil {
		t.Errorf("expected error but no errors ouccured")
	}
}

func TestApiPost(t *testing.T) {
	cases := map[string]struct {
		client               *http.Client
		expectHasError       bool
		timeout              int
		expectedErrorMessage string
		expectedText         string
	}{
		"normal": {
			client:         client(t, 0, nil),
			expectHasError: false,
			timeout:        10000,
			expectedText:   "{\"Status\":\"OK\"}",
		},
		"204NoContent": {
			client: client(t, 0, &http.Response{
				StatusCode: http.StatusNoContent,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("{\"Status\":\"OK\"}"))),
				Header:     make(http.Header),
			}),
			expectHasError: false,
			timeout:        10000,
			expectedText:   "{\"Status\":\"OK\"}",
		},
		"timeout": {
			client:               client(t, 0, nil),
			expectHasError:       true,
			timeout:              0,
			expectedErrorMessage: "HTTP request cancelled",
		},
		"500Error": {
			client: client(t, 0, &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("error"))),
				Header:     make(http.Header),
			}),
			expectHasError:       true,
			timeout:              10000,
			expectedErrorMessage: "HTTP Client error!: 500",
		},
		"400Error": {
			client: client(t, 0, &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("error"))),
				Header:     make(http.Header),
			}),
			expectHasError:       true,
			timeout:              10000,
			expectedErrorMessage: "HTTP Client error!: 400",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			e := utils.NewHttpClient(c.timeout, utils.OptionHTTPClient(c.client))

			resp, err := e.Post("https://example.com", "application/json", []byte("test"))

			if c.expectHasError {
				if err == nil {
					t.Errorf("expected error but no errors ouccured")
					return
				}

				if err.Error() != c.expectedErrorMessage {
					t.Errorf("unexpected error message. expected '%s', actual '%s'", c.expectedErrorMessage, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf(err.Error())
				return
			}

			if string(resp) != c.expectedText {
				t.Errorf("unexpected response's text. expected '%s', actual '%s'", c.expectedText, string(resp))
			}
		})
	}
}

func TestApiPut(t *testing.T) {
	cases := map[string]struct {
		client               *http.Client
		expectHasError       bool
		timeout              int
		expectedErrorMessage string
		expectedText         string
	}{
		"normal": {
			client:         client(t, 0, nil),
			expectHasError: false,
			timeout:        10000,
			expectedText:   "{\"Status\":\"OK\"}",
		},
		"204NoContent": {
			client: client(t, 0, &http.Response{
				StatusCode: http.StatusNoContent,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("{\"Status\":\"OK\"}"))),
				Header:     make(http.Header),
			}),
			expectHasError: false,
			timeout:        10000,
			expectedText:   "{\"Status\":\"OK\"}",
		},
		"timeout": {
			client:               client(t, 0, nil),
			expectHasError:       true,
			timeout:              0,
			expectedErrorMessage: "HTTP request cancelled",
		},
		"500Error": {
			client: client(t, 0, &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("error"))),
				Header:     make(http.Header),
			}),
			expectHasError:       true,
			timeout:              10000,
			expectedErrorMessage: "HTTP Client error!: 500",
		},
		"400Error": {
			client: client(t, 0, &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("error"))),
				Header:     make(http.Header),
			}),
			expectHasError:       true,
			timeout:              10000,
			expectedErrorMessage: "HTTP Client error!: 400",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			e := utils.NewHttpClient(c.timeout, utils.OptionHTTPClient(c.client))

			err := e.Put("https://example.com", "application/json", []byte("test"))

			if c.expectHasError {
				if err == nil {
					t.Errorf("expected error but no errors ouccured")
					return
				}

				if err.Error() != c.expectedErrorMessage {
					t.Errorf("unexpected error message. expected '%s', actual '%s'", c.expectedErrorMessage, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf(err.Error())
				return
			}
		})
	}
}

func TestApiDelete(t *testing.T) {
	cases := map[string]struct {
		client               *http.Client
		expectHasError       bool
		timeout              int
		expectedErrorMessage string
		expectedText         string
	}{
		"normal": {
			client:         client(t, 0, nil),
			expectHasError: false,
			timeout:        10000,
			expectedText:   "{\"Status\":\"OK\"}",
		},
		"204NoContent": {
			client: client(t, 0, &http.Response{
				StatusCode: http.StatusNoContent,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("{\"Status\":\"OK\"}"))),
				Header:     make(http.Header),
			}),
			expectHasError: false,
			timeout:        10000,
			expectedText:   "{\"Status\":\"OK\"}",
		},
		"timeout": {
			client:               client(t, 0, nil),
			expectHasError:       true,
			timeout:              0,
			expectedErrorMessage: "HTTP request cancelled",
		},
		"500Error": {
			client: client(t, 0, &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("error"))),
				Header:     make(http.Header),
			}),
			expectHasError:       true,
			timeout:              10000,
			expectedErrorMessage: "HTTP Client error!: 500",
		},
		"400Error": {
			client: client(t, 0, &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("error"))),
				Header:     make(http.Header),
			}),
			expectHasError:       true,
			timeout:              10000,
			expectedErrorMessage: "HTTP Client error!: 400",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			e := utils.NewHttpClient(c.timeout, utils.OptionHTTPClient(c.client))

			err := e.Delete("https://example.com")

			if c.expectHasError {
				if err == nil {
					t.Errorf("expected error but no errors ouccured")
					return
				}

				if err.Error() != c.expectedErrorMessage {
					t.Errorf("unexpected error message. expected '%s', actual '%s'", c.expectedErrorMessage, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf(err.Error())
				return
			}
		})
	}
}

func TestNewProxyClient(t *testing.T) { //しょぼいテスト...
	actual, err := utils.NewProxyClient("http://example.com")
	if err != nil {
		t.Errorf("Error occurred: '%s'", err)
	}
	if reflect.TypeOf(actual).Kind().String() != "ptr" {
		t.Errorf("Types not match '%s'", reflect.TypeOf(actual).Kind().String())
	}
}

func ExampleApi_Get() {
	client := utils.NewHttpClient(10000)
	resp, _ := client.Get("https://aws-health-dashboard.vercel.app/api/aws")
	fmt.Println(string(resp))
}

func ExampleApi_Post() {
	client := utils.NewHttpClient(10000)
	resp, _ := client.Post("https://aws-health-dashboard.vercel.app/api/aws", "application/json", []byte("{\"test\": \"test\"}"))
	fmt.Println(string(resp))
}

func ExampleApi_Put() {
	client := utils.NewHttpClient(10000)
	_ = client.Put("https://example.com/test", "application/json", []byte("{\"test\": \"test\"}"))
}

func ExampleApi_Delete() {
	client := utils.NewHttpClient(10000)
	_ = client.Delete("https://example.com/test")
}

func ExampleNewProxyClient() {
	proxy, _ := utils.NewProxyClient("http://tubone24:password@127.0.0.1:8080")
	client := utils.NewHttpClient(10000,  utils.OptionHTTPClient(proxy))
	// via Proxy
	resp, _ := client.Get("https://aws-health-dashboard.vercel.app/api/aws")
	fmt.Println(string(resp))
}
