package utils_test

import (
	"bytes"
	"github.com/tubone24/redump/pkg/utils"
	"io/ioutil"
	"net/http"
	"testing"
	"github.com/goccy/go-json"
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

func TestApi_Get(t *testing.T) {
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
			timeout: 10000,
			expectedText:   "{\"Status\":\"OK\"}",
		},
		"204NoContent": {
			client:         client(t, 0, &http.Response{
				StatusCode: http.StatusNoContent,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("{\"Status\":\"OK\"}"))),
				Header:     make(http.Header),
			}),
			expectHasError: false,
			timeout: 10000,
			expectedText:   "{\"Status\":\"OK\"}",
		},
		"timeout": {
			client:         client(t, 0, nil),
			expectHasError: true,
			timeout: 0,
			expectedErrorMessage: "HTTP request cancelled",
		},
		"500Error": {
			client:         client(t, 0, &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("error"))),
				Header:     make(http.Header),
			}),
			expectHasError: true,
			timeout: 10000,
			expectedErrorMessage: "HTTP Client error!: 500",
		},
		"400Error": {
			client:         client(t, 0, &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("error"))),
				Header:     make(http.Header),
			}),
			expectHasError: true,
			timeout: 10000,
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