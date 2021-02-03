package redmine_test

import (
	"bytes"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/redmine"
	"io/ioutil"
	"net/http"
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

var issueJson = redmine.Issue{
	Id:      1,
	Project: redmine.Project{Id: 1, Name: "testProject"},
	Tracker: redmine.Tracker{Id: 1, Name: "doing"},
	Status:  redmine.Status{Id: 1, Name: "test"}, Priority: redmine.Priority{Id: 1, Name: "High"},
	Author:      redmine.Author{Id: 1, Name: "testUser"},
	AssignedTo:  redmine.AssignedTo{Id: 1, Name: "testUser"},
	Subject:     "test1",
	Description: "testtesttesttest",
	StartDate:   "2020-01-01T00:00:00Z",
	CustomFields: redmine.CustomFields{&redmine.CustomField{
		Id:       1,
		Name:     "customField1",
		Multiple: true,
		Value:    []string{"aaaa", "bbb", "ccc"}}},
	CreatedOn: "2020-01-01T00:00:00Z",
	UpdatedOn: "2020-01-01T00:00:00Z",
	Attachments: redmine.Attachments{&redmine.Attachment{
		Id: 1, FileName: "test.png",
		FileSize:    12000,
		ContentUrl:  "http://example.com/test.png",
		Description: "testFile",
		Author:      redmine.Author{Id: 1, Name: "testUser"},
		CreatedOn:   "2020-01-01T00:00:00Z"}},
	Journals: redmine.Journals{&redmine.Journal{
		Id:        1,
		User:      redmine.User{Id: 1, Name: "testUser"},
		Notes:     "testtest",
		CreatedOn: "2020-01-01T00:00:00Z"},
		&redmine.Journal{
			Id:        2,
			User:      redmine.User{Id: 1, Name: "testUser"},
			Notes:     "testtest2",
			CreatedOn: "2020-01-01T00:00:00Z"},
		&redmine.Journal{
			Id:        3,
			User:      redmine.User{Id: 1, Name: "testUser"},
			Notes:     "testtest",
			CreatedOn: "2020-01-01T00:00:00Z", Details: redmine.Details{&redmine.Detail{
				Property: "change",
				Name:     "upload",
				OldValue: "aaa",
				NewValue: "bbb"}},},
	},
	Watchers: redmine.Watchers{&redmine.Watcher{
		Id: 1, Name: "testUser"}, &redmine.Watcher{Id: 2, Name: "testUser2"}, &redmine.Watcher{Id: 3, Name: "testUser3"},},
}

func clientIssues(t *testing.T, respTime time.Duration, resp *http.Response) *http.Client {
	var issuesResult struct {
		Issues     redmine.Issues `json:"issues"`
		TotalCount int            `json:"total_count"`
		Offset     int            `json:"offset"`
		Limit      int            `json:"limit"`
	}
	issuesResult.Issues = redmine.Issues{}
	issuesResult.TotalCount = 10000
	issuesResult.Limit = 100
	issuesResult.Offset = 0
	for i := 0; i < 100; i++ {
		issuesResult.Issues = append(issuesResult.Issues, &issueJson)
	}
	t.Helper()

	b, err := json.Marshal(&issuesResult)
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

func clientIssue(t *testing.T, respTime time.Duration, resp *http.Response) *http.Client {
	var issueResult struct {
		Issue redmine.Issue `json:"issue"`
	}
	issueResult.Issue = issueJson
	t.Helper()

	b, err := json.Marshal(&issueResult)
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

func clientDownloadAttachments(t *testing.T, respTime time.Duration, resp *http.Response) *http.Client {
	b := []byte("test")

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

func TestGetIssues(t *testing.T) {
	resp, err := redmine.GetIssues("https://example.com", "aaa", 0, 10000, clientIssues(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if resp[0].Id != issueJson.Id {
		t.Errorf("expected: %d, actual %d", resp[0].Id, issueJson.Id)
	}
}

func TestGetIssue(t *testing.T) {
	resp, err := redmine.GetIssues("https://example.com", "aaa", 0, 10000, clientIssue(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if resp[0].Id != issueJson.Id {
		t.Errorf("expected: %d, actual %d", resp[0].Id, issueJson.Id)
	}
}

func TestDownloadAttachmentFiles(t *testing.T) {
	resp, err := redmine.DownloadAttachmentFiles("aaa", 10000, issueJson.Attachments, clientDownloadAttachments(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if string(resp[0]) != "test" {
		t.Errorf("expected: %s, actual %s", string(resp[0]), "test")
	}
}
