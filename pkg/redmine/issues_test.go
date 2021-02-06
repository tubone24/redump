package redmine_test

import (
	"bytes"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/redmine"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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

var uploadFile = redmine.FileParam{
	FileName: "test.png",
	ContentType: "image/png",
	Contents: []byte{},
	Token: "tokentoken",
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

var issueJsonNoAttachments = redmine.Issue{
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
	Attachments: nil,
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

var issueJsonParam = redmine.IssueParam{
	ProjectId: 1,
	TrackerId: 1,
	StatusId: 1,
	PriorityId: 1,
	Subject: "test1",
	Description: "testtesttesttest",
	AssignedToId: 1,
	CustomFields: redmine.CustomFields{
		&redmine.CustomField{Id: 1},
	},
	Notes: "testNote",
	Uploads: []redmine.Uploads{
		redmine.Uploads{
			Token: "aaa",
			FileName: "aaa.png",
			ContentType: "image/png",
		},
	},
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
		t.Helper()
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
	resp, err := redmine.GetIssue("https://example.com", "aaa", 1, 10000, clientIssue(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if resp.Id != issueJson.Id {
		t.Errorf("expected: %d, actual %d", resp.Id, issueJson.Id)
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

func TestCreateIssueFromByteSlice(t *testing.T) {
	const bufferSize = 256
	var content []byte
	dir, _ := os.Getwd()
	filename := filepath.FromSlash(dir + "/../../tests/test_assets/2463.json")
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	buffer := make([]byte, bufferSize)
	for {
		n, err := fp.Read(buffer)
		if 0 < n {
			content = append(content, buffer...)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Errorf("Error occured: %s", err)
		}
	}
	resp, err := redmine.CreateIssueFromByteSlice(content)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if resp.Id != 2463 {
		t.Errorf("expected: %d, actual %d", 2463, resp.Id)
	}
}

func TestCreateIssueParam(t *testing.T) {
	var uploadFiles []redmine.FileParam
	uploadFiles = append(uploadFiles, uploadFile)
	issueParam := redmine.CreateIssueParam(issueJson, uploadFiles)
	if issueParam.ProjectId != issueJson.Project.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Project.Id, issueParam.ProjectId)
	}
	if issueParam.TrackerId != issueJson.Tracker.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Tracker.Id, issueParam.TrackerId)
	}
	if issueParam.StatusId != issueJson.Status.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Status.Id, issueParam.StatusId)
	}
	if issueParam.PriorityId != issueJson.Priority.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Priority.Id, issueParam.PriorityId)
	}
	if issueParam.Subject != issueJson.Subject {
		t.Errorf("expected: %s, actual %s", issueJson.Subject, issueParam.Subject)
	}
	if issueParam.Description != issueJson.Description {
		t.Errorf("expected: %s, actual %s", issueJson.Subject, issueParam.Subject)
	}
	if issueParam.AssignedToId != issueJson.AssignedTo.Id {
		t.Errorf("expected: %d, actual %d", issueJson.AssignedTo.Id, issueParam.AssignedToId)
	}
	if issueParam.CustomFields[0].Id != issueJson.CustomFields[0].Id {
		t.Errorf("expected: %d, actual %d", issueJson.CustomFields[0].Id, issueParam.CustomFields[0].Id)
	}
	if issueParam.Uploads[0].FileName != issueJson.Attachments[0].FileName {
		t.Errorf("expected: %s, actual %s", issueJson.Attachments[0].FileName, issueParam.Uploads[0].FileName)
	}
}

func TestCreateIssueParamNoFile(t *testing.T) {
	var uploadFiles []redmine.FileParam
	issueParam := redmine.CreateIssueParam(issueJsonNoAttachments, uploadFiles)
	if issueParam.ProjectId != issueJson.Project.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Project.Id, issueParam.ProjectId)
	}
	if issueParam.TrackerId != issueJson.Tracker.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Tracker.Id, issueParam.TrackerId)
	}
	if issueParam.StatusId != issueJson.Status.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Status.Id, issueParam.StatusId)
	}
	if issueParam.PriorityId != issueJson.Priority.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Priority.Id, issueParam.PriorityId)
	}
	if issueParam.Subject != issueJson.Subject {
		t.Errorf("expected: %s, actual %s", issueJson.Subject, issueParam.Subject)
	}
	if issueParam.Description != issueJson.Description {
		t.Errorf("expected: %s, actual %s", issueJson.Subject, issueParam.Subject)
	}
	if issueParam.AssignedToId != issueJson.AssignedTo.Id {
		t.Errorf("expected: %d, actual %d", issueJson.AssignedTo.Id, issueParam.AssignedToId)
	}
	if issueParam.CustomFields[0].Id != issueJson.CustomFields[0].Id {
		t.Errorf("expected: %d, actual %d", issueJson.CustomFields[0].Id, issueParam.CustomFields[0].Id)
	}
	if issueParam.Uploads != nil {
		t.Errorf("expected: nil, actual %s", issueParam.Uploads)
	}
}

func TestCreateIssue(t *testing.T) {
	issueId, err := redmine.CreateIssue("http://example.com", "xxxx", 10000, issueJsonParam, clientIssue(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if issueId != issueJson.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Id, issueId)
	}
}

func TestCreateJournalStrings(t *testing.T) {
	notes := redmine.CreateJournalStrings(issueJson)
	if notes[0] != issueJson.Journals[0].Notes {
		t.Errorf("expected: %s, actual %s", issueJson.Journals[0].Notes, notes[0])
	}
}

func TestUpdateIssueJournals(t *testing.T) {
	journals := []string{"test1", "test2", "test3"}
	err := redmine.UpdateIssueJournals("http://example.com", "xxxxx", 1, 10000, journals, clientIssues(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
}

func TestDeleteIssue(t *testing.T) {
	err := redmine.DeleteIssue("http://examole.com", "xxxxx", 1, 10000, clientIssues(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
}

func TestUpdateWatchers(t *testing.T) {
	err := redmine.UpdateWatchers("http://example.com", "xxxxx", 1, 10000, issueJson, clientIssues(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
}

func TestUploadAttachmentFiles(t *testing.T) {
	var uploadFiles []redmine.FileParam
	uploadFiles = append(uploadFiles, uploadFile)
	resp, err := redmine.UploadAttachmentFiles("http://example.com", "xxxxx", 10000, uploadFiles, clientIssues(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if resp[0].FileName != uploadFile.FileName {
		t.Errorf("expected: %s, actual %s", uploadFile.FileName, resp[0].FileName)
	}
}

func TestUnmarshalByteIssue(t *testing.T) {
	resp, err := json.Marshal(&issueJson)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	actual, err := redmine.UnmarshalByteIssue(resp)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if actual.Id != issueJson.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Id, actual.Id)
	}
}
