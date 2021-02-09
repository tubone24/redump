package redmine_test

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/tubone24/redump/pkg/utils"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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
	FileName:    "test.png",
	ContentType: "image/png",
	Contents:    []byte{},
	Token:       "tokentoken",
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
				NewValue: "bbb"}}},
	},
	Watchers: redmine.Watchers{&redmine.Watcher{
		Id: 1, Name: "testUser"}, &redmine.Watcher{Id: 2, Name: "testUser2"}, &redmine.Watcher{Id: 3, Name: "testUser3"}},
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
	CreatedOn:   "2020-01-01T00:00:00Z",
	UpdatedOn:   "2020-01-01T00:00:00Z",
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
				NewValue: "bbb"}}},
	},
	Watchers: redmine.Watchers{&redmine.Watcher{
		Id: 1, Name: "testUser"}, &redmine.Watcher{Id: 2, Name: "testUser2"}, &redmine.Watcher{Id: 3, Name: "testUser3"}},
}

var issueJsonInvalidAttachmentsUrl = redmine.Issue{
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
		ContentUrl:  "http://examplaspmaf.asdfof.qfqwfpkfw.qwffwqfoffmp,wfpi09mopf12rjqsffcaf.fasfkfoafasfakfa,fa.fasf.fasf.sfa.fasf.afe.com.com.com.jp.com.jp/test.png",
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
				NewValue: "bbb"}}},
	},
	Watchers: redmine.Watchers{&redmine.Watcher{
		Id: 1, Name: "testUser"}, &redmine.Watcher{Id: 2, Name: "testUser2"}, &redmine.Watcher{Id: 3, Name: "testUser3"}},
}

var issueJsonParam = redmine.IssueParam{
	ProjectId:    1,
	TrackerId:    1,
	StatusId:     1,
	PriorityId:   1,
	Subject:      "test1",
	Description:  "testtesttesttest",
	AssignedToId: 1,
	CustomFields: redmine.CustomFields{
		&redmine.CustomField{Id: 1},
	},
	Notes: "testNote",
	Uploads: []redmine.Uploads{
		redmine.Uploads{
			Token:       "aaa",
			FileName:    "aaa.png",
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
		var ij = redmine.Issue{
			Id:      i,
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
						NewValue: "bbb"}}},
			},
			Watchers: redmine.Watchers{&redmine.Watcher{
				Id: 1, Name: "testUser"}, &redmine.Watcher{Id: 2, Name: "testUser2"}, &redmine.Watcher{Id: 3, Name: "testUser3"}},
		}
		issuesResult.Issues = append(issuesResult.Issues, &ij)
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

func clientIssuesInvalidResp(t *testing.T, respTime time.Duration, resp *http.Response) *http.Client {
	t.Helper()
	return NewTestClient(func(req *http.Request) *http.Response {
		time.Sleep(respTime)
		if resp != nil {
			return resp
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
			Header:     make(http.Header),
		}
	})
}

func clientIssuesBench(b *testing.B, respTime time.Duration, resp *http.Response, issues redmine.Issues) *http.Client {
	var issuesResult struct {
		Issues     redmine.Issues `json:"issues"`
		TotalCount int            `json:"total_count"`
		Offset     int            `json:"offset"`
		Limit      int            `json:"limit"`
	}
	issuesResult.Issues = redmine.Issues{}
	issuesResult.TotalCount = 100000
	issuesResult.Limit = 100
	issuesResult.Offset = 0
	issuesResult.Issues = issues
	b.Helper()

	bu, err := json.Marshal(&issuesResult)
	if err != nil {
		b.Fatal(err)
	}

	return NewTestClient(func(req *http.Request) *http.Response {
		time.Sleep(respTime)
		if resp != nil {
			return resp
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(bu)),
			Header:     make(http.Header),
		}
	})
}

func clientIssueBench(b *testing.B, respTime time.Duration, resp *http.Response) *http.Client {
	var issueResult struct {
		Issue redmine.Issue `json:"issue"`
	}
	issueResult.Issue = issueJson
	b.Helper()

	bu, err := json.Marshal(&issueResult)
	if err != nil {
		b.Fatal(err)
	}

	return NewTestClient(func(req *http.Request) *http.Response {
		time.Sleep(respTime)
		if resp != nil {
			return resp
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(bu)),
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

func clientDownloadAttachmentsBench(b *testing.B, respTime time.Duration, resp *http.Response) *http.Client {
	bu, _ := json.Marshal(issueJson)

	return NewTestClient(func(req *http.Request) *http.Response {
		time.Sleep(respTime)
		if resp != nil {
			return resp
		}
		b.Helper()
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(bu)),
			Header:     make(http.Header),
		}
	})
}

func TestGetIssues(t *testing.T) {
	resp, err := redmine.GetIssues("https://example.com", "aaa", 0, 10000, clientIssues(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if resp[0].Id != 0 {
		t.Errorf("expected: %d, actual %d", resp[0].Id, 0)
	}
	if resp[99].Id != 99 {
		t.Errorf("expected: %d, actual %d", resp[0].Id, 99)
	}
	if resp[100].Id != 0 {
		t.Errorf("expected: %d, actual %d", resp[0].Id, 0)
	}
	if resp[101].Id != 1 {
		t.Errorf("expected: %d, actual %d", resp[0].Id, 1)
	}
}

func TestGetIssuesWithProjectId(t *testing.T) {
	resp, err := redmine.GetIssues("https://example.com", "aaa", 1, 10000, clientIssues(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if resp[0].Id != 0 {
		t.Errorf("expected: %d, actual %d", resp[0].Id, 0)
	}
}

func TestGetIssuesInvalidUrl(t *testing.T) {
	_, err := redmine.GetIssues("https://examphoiaiho.dq.dq.hfnqwopfq.c,cpckckpc.kdjow-wq-d-qdqd-qdd-qdqdccwsccl.le.com.jkp.laala.com.com.com", "aaa", 1, 1000, nil)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestGetIssuesInvalidResp(t *testing.T) {
	_, err := redmine.GetIssues("https://example.com", "aaa", 0, 10000, clientIssuesInvalidResp(t, 1000, nil))
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestGetIssue(t *testing.T) {
	resp, err := redmine.GetIssue("https://example.com", "aaa", 1, 10000, clientIssue(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if resp.Id != issueJson.Id {
		t.Errorf("expected: %d, actual %d", resp.Id, issueJson.Id)
	}
}

func TestGetIssueInvalidUrl(t *testing.T) {
	_, err := redmine.GetIssue("https://examphoiaiho.dq.dq.hfnqwopfq.c,cpckckpc.kdjow-wq-d-qdqd-qdd-qdqdccwsccl.le.com.jkp.laala.com.com.com", "aaa", 1, 1000, nil)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestGetIssueInvalidResp(t *testing.T) {
	_, err := redmine.GetIssue("https://example.com", "aaa", 1, 1000, clientIssuesInvalidResp(t, 1000, nil))
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestDownloadAttachmentFiles(t *testing.T) {
	resp, err := redmine.DownloadAttachmentFiles("aaa", 10000, issueJson.Attachments, clientDownloadAttachments(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if string(resp[0]) != "test" {
		t.Errorf("expected: %s, actual %s", string(resp[0]), "test")
	}
}

func TestDownloadAttachmentFilesInvalidUrl(t *testing.T) {
	_, err := redmine.DownloadAttachmentFiles("aaa", 1000, issueJsonInvalidAttachmentsUrl.Attachments, nil)
	if err == nil {
		t.Errorf("Error not occurred")
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
			t.Errorf("Error occurred: %s", err)
		}
	}
	resp, err := redmine.CreateIssueFromByteSlice(content)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if resp.Id != 2463 {
		t.Errorf("expected: %d, actual %d", 2463, resp.Id)
	}
}

func TestCreateIssueFromByteSliceInvalidBytes(t *testing.T) {
	_, err := redmine.CreateIssueFromByteSlice([]byte("aaa"))
	if err == nil {
		t.Errorf("Error not occurred")
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
		t.Errorf("Error occurred: %s", err)
	}
	if issueId != issueJson.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Id, issueId)
	}
}

func TestCreateIssueInvalidUrl(t *testing.T) {
	_, err := redmine.CreateIssue("https://examphoiaiho.dq.dq.hfnqwopfq.c,cpckckpc.kdjow-wq-d-qdqd-qdd-qdqdccwsccl.le.com.jkp.laala.com.com.com", "xxxx", 1000, issueJsonParam, nil)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestCreateIssueInvalidResp(t *testing.T) {
	_, err := redmine.CreateIssue("https://example.com", "xxxx", 1000, issueJsonParam, clientIssuesInvalidResp(t, 1000, nil))
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestCreateJournalStrings(t *testing.T) {
	notes := redmine.CreateJournalStrings(issueJson)
	if notes[0] != issueJson.Journals[0].Notes {
		t.Errorf("expected: %s, actual %s", issueJson.Journals[0].Notes, notes[0])
	}
}

func TestUpdateIssueJournals(t *testing.T) {
	journals := []string{"test1", "test2", "test3", "", "", "test4", "test5"}
	err := redmine.UpdateIssueJournals("http://example.com", "xxxxx", 1, 10000, journals, clientIssues(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
}

func TestUpdateIssueJournalsInvalidUrl(t *testing.T) {
	journals := []string{"test1", "test2", "test3", "", "", "test4", "test5"}
	err := redmine.UpdateIssueJournals("https://examphoiaiho.dq.dq.hfnqwopfq.c,cpckckpc.kdjow-wq-d-qdqd-qdd-qdqdccwsccl.le.com.jkp.laala.com.com.com", "xxxxx", 1, 10000, journals, nil)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestDeleteIssueInvalidUrl(t *testing.T) {
	err := redmine.DeleteIssue("https://examphoiaiho.dq.dq.hfnqwopfq.c,cpckckpc.kdjow-wq-d-qdqd-qdd-qdqdccwsccl.le.com.jkp.laala.com.com.com", "xxxxx", 1, 10000, nil)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestDeleteIssue(t *testing.T) {
	err := redmine.DeleteIssue("http://examole.com", "xxxxx", 1, 10000, clientIssues(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
}

func TestUpdateWatchers(t *testing.T) {
	err := redmine.UpdateWatchers("http://example.com", "xxxxx", 1, 10000, issueJson, clientIssue(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
}

func TestUpdateWatchersInvalidUrl(t *testing.T) {
	err := redmine.UpdateWatchers("https://examphoiaiho.dq.dq.hfnqwopfq.c,cpckckpc.kdjow-wq-d-qdqd-qdd-qdqdccwsccl.le.com.jkp.laala.com.com.com", "xxxxx", 1, 1000, issueJson, nil)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestUploadAttachmentFiles(t *testing.T) {
	var uploadFiles []redmine.FileParam
	uploadFiles = append(uploadFiles, uploadFile)
	resp, err := redmine.UploadAttachmentFiles("http://example.com", "xxxxx", 10000, uploadFiles, clientIssue(t, 1000, nil))
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if resp[0].FileName != uploadFile.FileName {
		t.Errorf("expected: %s, actual %s", uploadFile.FileName, resp[0].FileName)
	}
}

func TestUploadAttachmentFilesInvalidUrl(t *testing.T) {
	var uploadFiles []redmine.FileParam
	uploadFiles = append(uploadFiles, uploadFile)
	_, err := redmine.UploadAttachmentFiles("https://examphoiaiho.dq.dq.hfnqwopfq.c,cpckckpc.kdjow-wq-d-qdqd-qdd-qdqdccwsccl.le.com.jkp.laala.com.com.com", "xxxxx", 1000, uploadFiles, nil)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestUploadAttachmentFilesInvalidBytes(t *testing.T) {
	var uploadFiles []redmine.FileParam
	uploadFiles = append(uploadFiles, uploadFile)
	_, err := redmine.UploadAttachmentFiles("http://example.com", "xxxxx", 10000, uploadFiles, clientIssuesInvalidResp(t, 1000, nil))
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestUnmarshalByteIssue(t *testing.T) {
	resp, err := json.Marshal(&issueJson)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	actual, err := redmine.UnmarshalByteIssue(resp)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if actual.Id != issueJson.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Id, actual.Id)
	}
}

func TestUnmarshalByteIssueInvalidBytes(t *testing.T) {
	_, err := redmine.UnmarshalByteIssue([]byte("{}"))
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func BenchmarkGetIssues(b *testing.B) {
	issueJsons100 := make(redmine.Issues, 100)
	for i := 0; i <100; i++ {
		ij := redmine.Issue{
			Id:      i,
			Project: redmine.Project{Id: 1, Name: "testProject"},
			Tracker: redmine.Tracker{Id: 1, Name: "doing"},
			Status:  redmine.Status{Id: 1, Name: "test"}, Priority: redmine.Priority{Id: 1, Name: "High"},
			Author:      redmine.Author{Id: 1, Name: "testUser"},
			AssignedTo:  redmine.AssignedTo{Id: 1, Name: "testUser"},
			Subject:     "test1" + strconv.Itoa(i),
			Description: "testtesttesttest" + strconv.Itoa(i),
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
						NewValue: "bbb"}}},
			},
			Watchers: redmine.Watchers{&redmine.Watcher{
				Id: 1, Name: "testUser"}, &redmine.Watcher{Id: 2, Name: "testUser2"}, &redmine.Watcher{Id: 3, Name: "testUser3"}},
		}
		issueJsons100[i] = &ij
	}
	b.ResetTimer()
	_, _ = redmine.GetIssue("https://example.com", "aaa", 1, 10000, clientIssuesBench(b, 0, nil, issueJsons100))
}

func BenchmarkGetIssue(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 100; i++ {
		_, _ = redmine.GetIssue("https://example.com", "aaa", i, 10000, clientIssueBench(b, 0, nil))
	}
}

func BenchmarkDownloadAttachmentFiles(b *testing.B) {
	bigAttachment := make(redmine.Attachments, 1000)
	for i := 0; i < 1000; i++ {
		at := redmine.Attachment{
			Id: 1, FileName: "test" + strconv.Itoa(i) + ".png",
			FileSize:    12000,
			ContentUrl:  "http://example.com/test" + strconv.Itoa(i) + ".png",
			Description: "testFile",
			Author:      redmine.Author{Id: 1, Name: "testUser"},
			CreatedOn:   "2020-01-01T00:00:00Z"}
		bigAttachment[i] = &at
	}
	b.ResetTimer()
	_, _ = redmine.DownloadAttachmentFiles("aaa", 10000, bigAttachment, clientDownloadAttachmentsBench(b, 0, nil))
}

func BenchmarkCreateIssueParam(b *testing.B) {
	uploadFiles := make([]redmine.FileParam, 1000)
	for i := 0; i < 1000; i++ {
		uf := redmine.FileParam{
			FileName:    "test" + strconv.Itoa(i) + ".png",
			ContentType: "image/png",
			Contents:    []byte("testtesttesttesttesttesttesttesttesttesttest"),
			Token:       "tokentoken" + strconv.Itoa(i),
		}
		uploadFiles[i] = uf
	}
	b.ResetTimer()
	_ = redmine.CreateIssueParam(issueJson, uploadFiles)
}

func BenchmarkUploadAttachmentFiles(b *testing.B) {
	uploadFiles := make([]redmine.FileParam, 1000)
	for i := 0; i < 1000; i++ {
		uf := redmine.FileParam{
			FileName:    "test" + strconv.Itoa(i) + ".png",
			ContentType: "image/png",
			Contents:    []byte("testtesttesttesttesttesttesttesttesttesttest"),
			Token:       "tokentoken" + strconv.Itoa(i),
		}
		uploadFiles[i] = uf
	}
	b.ResetTimer()
	_, _ = redmine.UploadAttachmentFiles("http://example.com", "xxxxx", 10000, uploadFiles, clientIssueBench(b, 0, nil))
}

func BenchmarkCreateJournalStrings(b *testing.B) {
	var bigIssueJson redmine.Issue
	journalsJson := make(redmine.Journals, 10000)
	for i := 0; i < 10000; i++ {
		journalsJson[i] = &redmine.Journal{Notes: "testtesttesttesttesttest"}
	}
	bigIssueJson.Journals = journalsJson
	b.ResetTimer()
	_ = redmine.CreateJournalStrings(bigIssueJson)
}

func BenchmarkCreateIssue(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 100; i++ {
		_, _ = redmine.CreateIssue("http://example.com", "xxxx", 10000, issueJsonParam, clientIssueBench(b, 0, nil))
	}
}

func BenchmarkCreateIssueFromByteSlice(b *testing.B) {
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
			b.Errorf("Error occurred: %s", err)
		}
	}
	b.ResetTimer()
	_, _ = redmine.CreateIssueFromByteSlice(content)
}

func BenchmarkDeleteIssue(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 100; i++ {
		_ = redmine.DeleteIssue("http://examole.com", "xxxxx", i, 10000, clientIssueBench(b, 0, nil))
	}
}

func BenchmarkUpdateIssueJournals(b *testing.B) {
	benchJournals := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		benchJournals[i] = "test journal!!"
	}
	b.ResetTimer()
	_ = redmine.UpdateIssueJournals("http://example.com", "xxxxx", 1, 10000, benchJournals, clientIssueBench(b, 0, nil))
}

func BenchmarkUpdateWatchers(b *testing.B) {
	bigWatcher := make(redmine.Watchers, 1000)
	for i := 0; i < 1000; i++ {
		watcher := redmine.Watcher{
			Id: i,
			Name: "testUser",
		}
		bigWatcher[i] = &watcher
	}
	var bigWatcherIssue redmine.Issue
	bigWatcherIssue.Watchers = bigWatcher
	b.ResetTimer()
	_ = redmine.UpdateWatchers("http://example.com", "xxxxx", 1, 10000, bigWatcherIssue, clientIssueBench(b, 0, nil))
}

func ExampleGetIssues() {
	// All Projects
	issues, _ := redmine.GetIssues("https://redmine.example.com", "your-api-key-1234567890", 0, 10000, nil)
	fmt.Printf("'%#v'", issues)
	// Several Project (ex. ID is 1)
	issues, _ = redmine.GetIssues("https://redmine.example.com", "your-api-key-1234567890", 1, 10000, nil)
	fmt.Printf("'%#v'", issues)
}

func ExampleGetIssue() {
	// Several Project (ex. ID is 1)
	issue, _ := redmine.GetIssue("https://redmine.example.com", "your-api-key-1234567890", 1, 10000, nil)
	fmt.Printf("'%#v'", issue)
}

func ExampleDownloadAttachmentFiles() {
	attachment := redmine.Attachments{&redmine.Attachment{
		Id: 1, FileName: "test.png",
		FileSize:    12000,
		ContentUrl:  "http://example.com/test.png",
		Description: "testFile",
		Author:      redmine.Author{Id: 1, Name: "testUser"},
		CreatedOn:   "2020-01-01T00:00:00Z"}}

	resp, _ := redmine.DownloadAttachmentFiles("https://redmine.example.com", 10000, attachment, nil)
	// Return byte slice's slice. So for loop.
	for i, v := range resp {
		_ = utils.WriteFile("test"+strconv.Itoa(i)+".png", v)
	}
}

func ExampleCreateIssue() {
	var exampleParam = redmine.IssueParam{
		ProjectId:    1,
		TrackerId:    1,
		StatusId:     1,
		PriorityId:   1,
		Subject:      "test1",
		Description:  "testtesttesttest",
		AssignedToId: 1,
		CustomFields: redmine.CustomFields{
			&redmine.CustomField{Id: 1},
		},
		Notes: "testNote",
		Uploads: []redmine.Uploads{
			{
				Token:       "aaa",
				FileName:    "aaa.png",
				ContentType: "image/png",
			},
		},
	}
	resp, _ := redmine.CreateIssue("https://redmine.example.com", "your-api-key-1234567890", 10000, exampleParam, nil)
	// Issue ID
	fmt.Printf("'%d'", resp)
}

func ExampleCreateIssueFromByteSlice() {
	resp, _ := utils.ReadFile("test.json")
	issue, _ := redmine.CreateIssueFromByteSlice(resp)
	fmt.Printf("'%d'", issue.Id)
}

func ExampleCreateIssueParam() {
	var exampleIssue = redmine.Issue{
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
					NewValue: "bbb"}}},
		},
		Watchers: redmine.Watchers{&redmine.Watcher{
			Id: 1, Name: "testUser"}, &redmine.Watcher{Id: 2, Name: "testUser2"}, &redmine.Watcher{Id: 3, Name: "testUser3"}},
	}
	exampleUploadFiles := []redmine.FileParam{
		{
			FileName:    "test.png",
			ContentType: "image/png",
			Contents:    []byte{},
			Token:       "tokentoken",
		},
	}
	resp := redmine.CreateIssueParam(exampleIssue, exampleUploadFiles)
	fmt.Printf("'%#v'", resp)
}

func ExampleCreateJournalStrings() {
	var exampleIssue = redmine.Issue{
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
					NewValue: "bbb"}}},
		},
		Watchers: redmine.Watchers{&redmine.Watcher{
			Id: 1, Name: "testUser"}, &redmine.Watcher{Id: 2, Name: "testUser2"}, &redmine.Watcher{Id: 3, Name: "testUser3"}},
	}
	resp := redmine.CreateJournalStrings(exampleIssue)
	for _, v := range resp {
		fmt.Println(v)
	}
}

func ExampleUpdateIssueJournals() {
	exampleJournalsString := []string{"test", "test2", "test3"}
	_ = redmine.UpdateIssueJournals("https://redmine.example.com", "your-api-key-1234567890", 1, 10000, exampleJournalsString, nil)
}

func ExampleDeleteIssue() {
	_ = redmine.DeleteIssue("https://redmine.example.com", "your-api-key-1234567890", 1, 10000, nil)
}

func ExampleUpdateWatchers() {
	var exampleIssue = redmine.Issue{
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
					NewValue: "bbb"}}},
		},
		Watchers: redmine.Watchers{&redmine.Watcher{
			Id: 1, Name: "testUser"}, &redmine.Watcher{Id: 2, Name: "testUser2"}, &redmine.Watcher{Id: 3, Name: "testUser3"}},
	}
	_ = redmine.UpdateWatchers("https://redmine.example.com", "your-api-key-1234567890", 1, 10000, exampleIssue, nil)
}

func ExampleUploadAttachmentFiles() {
	ExampleuploadFile := []redmine.FileParam{
		{
			FileName:    "test.png",
			ContentType: "image/png",
			Contents:    []byte{},
		},
	}
	fileParams, _ := redmine.UploadAttachmentFiles("https://redmine.example.com", "your-api-key-1234567890", 10000, ExampleuploadFile, nil)
	fmt.Println(fileParams[0].Token)
}
