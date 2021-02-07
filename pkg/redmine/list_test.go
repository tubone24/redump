package redmine_test

import (
	"fmt"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/tubone24/redump/pkg/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestListProjectId(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "project_id.json")
	resp, err := redmine.ListProjectId(issuesJson, filename)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if resp[0].Id != issueJson.Project.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Project.Id, resp[0].Id)
	}
	if resp[0].Name != issueJson.Project.Name {
		t.Errorf("expected: %s, actual %s", issueJson.Project.Name, resp[0].Name)
	}
}

func TestListProjectIdInvalidFilename(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "project_id")
	fileDir2 := filepath.Join(filedir, "/project_id")
	err = utils.MakeDir(fileDir2)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	_, err = redmine.ListProjectId(issuesJson, filename)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestListTrackerId(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "tracker_id.json")
	resp, err := redmine.ListTrackerId(issuesJson, filename)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if resp[0].Id != issueJson.Tracker.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Tracker.Id, resp[0].Id)
	}
	if resp[0].Name != issueJson.Tracker.Name {
		t.Errorf("expected: %s, actual %s", issueJson.Tracker.Name, resp[0].Name)
	}
}

func TestListTrackerIdInvalidFilename(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "dummy")
	fileDir2 := filepath.Join(filedir, "/dummy")
	err = utils.MakeDir(fileDir2)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	_, err = redmine.ListTrackerId(issuesJson, filename)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestListStatusId(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "status_id.json")
	resp, err := redmine.ListStatusId(issuesJson, filename)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if resp[0].Id != issueJson.Status.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Status.Id, resp[0].Id)
	}
	if resp[0].Name != issueJson.Status.Name {
		t.Errorf("expected: %s, actual %s", issueJson.Status.Name, resp[0].Name)
	}
}

func TestListStatusIdInvalidFilename(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "dummy")
	fileDir2 := filepath.Join(filedir, "/dummy")
	err = utils.MakeDir(fileDir2)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	_, err = redmine.ListStatusId(issuesJson, filename)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestListPriorityId(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "priority_id.json")
	resp, err := redmine.ListPriorityId(issuesJson, filename)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if resp[0].Id != issueJson.Priority.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Priority.Id, resp[0].Id)
	}
	if resp[0].Name != issueJson.Priority.Name {
		t.Errorf("expected: %s, actual %s", issueJson.Priority.Name, resp[0].Name)
	}
}

func TestListPriorityIdInvalidFilename(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "dummy")
	fileDir2 := filepath.Join(filedir, "/dummy")
	err = utils.MakeDir(fileDir2)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	_, err = redmine.ListPriorityId(issuesJson, filename)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestListUserIdAssignedTo(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "user_id.json")
	resp, err := redmine.ListUserIdAssignedTo(issuesJson, filename)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if resp[0].Id != issueJson.AssignedTo.Id {
		t.Errorf("expected: %d, actual %d", issueJson.AssignedTo.Id, resp[0].Id)
	}
	if resp[0].Name != issueJson.AssignedTo.Name {
		t.Errorf("expected: %s, actual %s", issueJson.AssignedTo.Name, resp[0].Name)
	}
}

func TestListUserIdAssignedToInvalidFilename(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "dummy")
	fileDir2 := filepath.Join(filedir, "/dummy")
	err = utils.MakeDir(fileDir2)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	_, err = redmine.ListUserIdAssignedTo(issuesJson, filename)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestListCustomFieldsId(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "custom_fields_id_list.json")
	resp, err := redmine.ListCustomFieldsId(issuesJson, filename)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if resp[0].Id != issueJson.CustomFields[0].Id {
		t.Errorf("expected: %d, actual %d", issueJson.CustomFields[0].Id, resp[0].Id)
	}
	if resp[0].Name != issueJson.CustomFields[0].Name {
		t.Errorf("expected: %s, actual %s", issueJson.CustomFields[0].Name, resp[0].Name)
	}
}

func TestListCustomFieldsIdInvalidFilename(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "dummy")
	fileDir2 := filepath.Join(filedir, "/dummy")
	err = utils.MakeDir(fileDir2)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	_, err = redmine.ListCustomFieldsId(issuesJson, filename)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func ExampleListProjectId() {
	exampleIssues := redmine.Issues{{
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
	}}
	resp, _ := redmine.ListProjectId(exampleIssues, "")
	for _, v := range resp {
		fmt.Println(v.Name)
	}
}

func ExampleListTrackerId() {
	exampleIssues := redmine.Issues{{
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
	}}
	resp, _ := redmine.ListTrackerId(exampleIssues, "")
	for _, v := range resp {
		fmt.Println(v.Name)
	}
}

func ExampleListStatusId() {
	exampleIssues := redmine.Issues{{
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
	}}
	resp, _ := redmine.ListStatusId(exampleIssues, "")
	for _, v := range resp {
		fmt.Println(v.Name)
	}
}

func ExampleListPriorityId() {
	exampleIssues := redmine.Issues{{
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
	}}
	resp, _ := redmine.ListPriorityId(exampleIssues, "")
	for _, v := range resp {
		fmt.Println(v.Name)
	}
}

func ExampleListUserIdAssignedTo() {
	exampleIssues := redmine.Issues{{
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
	}}
	resp, _ := redmine.ListUserIdAssignedTo(exampleIssues, "")
	for _, v := range resp {
		fmt.Println(v.Name)
	}
}

func ExampleListCustomFieldsId() {
	exampleIssues := redmine.Issues{{
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
	}}
	resp, _ := redmine.ListCustomFieldsId(exampleIssues, "")
	for _, v := range resp {
		fmt.Println(v.Name)
	}
}
