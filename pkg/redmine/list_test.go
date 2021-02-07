package redmine_test

import (
	"fmt"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/tubone24/redump/pkg/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
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

func BenchmarkListProjectId(b *testing.B) {
	issueJsons10000 := make(redmine.Issues, 10000)
	for i := 0; i <10000; i++ {
		var ij redmine.Issue
		if i % 2 == 0 {
			ij = redmine.Issue{
				Id:      i,
				Project: redmine.Project{Id: i, Name: "testProject" + strconv.Itoa(i)},
				Tracker: redmine.Tracker{Id: i, Name: "doing" + strconv.Itoa(i)},
				Status:  redmine.Status{Id: i, Name: "test" + strconv.Itoa(i)}, Priority: redmine.Priority{Id: i, Name: "High" + strconv.Itoa(i)},
				Author:      redmine.Author{Id: i, Name: "testUser" + strconv.Itoa(i)},
				AssignedTo:  redmine.AssignedTo{Id: i, Name: "testUser" + strconv.Itoa(i)},
				Subject:     "test1" + strconv.Itoa(i),
				Description: "testtesttesttest" + strconv.Itoa(i),
				StartDate:   "2020-01-01T00:00:00Z",
				CustomFields: redmine.CustomFields{&redmine.CustomField{
					Id:       i,
					Name:     "customField" + strconv.Itoa(i),
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
		} else {
			ij = redmine.Issue{
				Id:      i,
				Project: redmine.Project{Id: i + 1, Name: "testProject" + strconv.Itoa(i + 1)},
				Tracker: redmine.Tracker{Id: i + 1, Name: "doing" + strconv.Itoa(i + 1)},
				Status:  redmine.Status{Id: i + 1, Name: "test" + strconv.Itoa(i + 1)}, Priority: redmine.Priority{Id: i + 1, Name: "High" + strconv.Itoa(i + 1)},
				Author:      redmine.Author{Id: i + 1, Name: "testUser" + strconv.Itoa(i + 1)},
				AssignedTo:  redmine.AssignedTo{Id: i + 1, Name: "testUser" + strconv.Itoa(i + 1)},
				Subject:     "test" + strconv.Itoa(i),
				Description: "testtesttesttest" + strconv.Itoa(i),
				StartDate:   "2020-01-01T00:00:00Z",
				CustomFields: redmine.CustomFields{&redmine.CustomField{
					Id:       i + 1,
					Name:     "customField" + strconv.Itoa(i + 1),
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
		}
		issueJsons10000[i] = &ij
	}
	_, _ = redmine.ListProjectId(issueJsons10000, "")
}

func BenchmarkListTrackerId(b *testing.B) {
	issueJsons10000 := make(redmine.Issues, 10000)
	for i := 0; i <10000; i++ {
		var ij redmine.Issue
		if i % 2 == 0 {
			ij = redmine.Issue{
				Id:      i,
				Project: redmine.Project{Id: i, Name: "testProject" + strconv.Itoa(i)},
				Tracker: redmine.Tracker{Id: i, Name: "doing" + strconv.Itoa(i)},
				Status:  redmine.Status{Id: i, Name: "test" + strconv.Itoa(i)}, Priority: redmine.Priority{Id: i, Name: "High" + strconv.Itoa(i)},
				Author:      redmine.Author{Id: i, Name: "testUser" + strconv.Itoa(i)},
				AssignedTo:  redmine.AssignedTo{Id: i, Name: "testUser" + strconv.Itoa(i)},
				Subject:     "test1" + strconv.Itoa(i),
				Description: "testtesttesttest" + strconv.Itoa(i),
				StartDate:   "2020-01-01T00:00:00Z",
				CustomFields: redmine.CustomFields{&redmine.CustomField{
					Id:       i,
					Name:     "customField" + strconv.Itoa(i),
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
		} else {
			ij = redmine.Issue{
				Id:      i,
				Project: redmine.Project{Id: i + 1, Name: "testProject" + strconv.Itoa(i + 1)},
				Tracker: redmine.Tracker{Id: i + 1, Name: "doing" + strconv.Itoa(i + 1)},
				Status:  redmine.Status{Id: i + 1, Name: "test" + strconv.Itoa(i + 1)}, Priority: redmine.Priority{Id: i + 1, Name: "High" + strconv.Itoa(i + 1)},
				Author:      redmine.Author{Id: i + 1, Name: "testUser" + strconv.Itoa(i + 1)},
				AssignedTo:  redmine.AssignedTo{Id: i + 1, Name: "testUser" + strconv.Itoa(i + 1)},
				Subject:     "test" + strconv.Itoa(i),
				Description: "testtesttesttest" + strconv.Itoa(i),
				StartDate:   "2020-01-01T00:00:00Z",
				CustomFields: redmine.CustomFields{&redmine.CustomField{
					Id:       i + 1,
					Name:     "customField" + strconv.Itoa(i + 1),
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
		}
		issueJsons10000[i] = &ij
	}
	_, _ = redmine.ListTrackerId(issueJsons10000, "")
}

func BenchmarkListStatusId(b *testing.B) {
	issueJsons10000 := make(redmine.Issues, 10000)
	for i := 0; i <10000; i++ {
		var ij redmine.Issue
		if i % 2 == 0 {
			ij = redmine.Issue{
				Id:      i,
				Project: redmine.Project{Id: i, Name: "testProject" + strconv.Itoa(i)},
				Tracker: redmine.Tracker{Id: i, Name: "doing" + strconv.Itoa(i)},
				Status:  redmine.Status{Id: i, Name: "test" + strconv.Itoa(i)}, Priority: redmine.Priority{Id: i, Name: "High" + strconv.Itoa(i)},
				Author:      redmine.Author{Id: i, Name: "testUser" + strconv.Itoa(i)},
				AssignedTo:  redmine.AssignedTo{Id: i, Name: "testUser" + strconv.Itoa(i)},
				Subject:     "test1" + strconv.Itoa(i),
				Description: "testtesttesttest" + strconv.Itoa(i),
				StartDate:   "2020-01-01T00:00:00Z",
				CustomFields: redmine.CustomFields{&redmine.CustomField{
					Id:       i,
					Name:     "customField" + strconv.Itoa(i),
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
		} else {
			ij = redmine.Issue{
				Id:      i,
				Project: redmine.Project{Id: i + 1, Name: "testProject" + strconv.Itoa(i + 1)},
				Tracker: redmine.Tracker{Id: i + 1, Name: "doing" + strconv.Itoa(i + 1)},
				Status:  redmine.Status{Id: i + 1, Name: "test" + strconv.Itoa(i + 1)}, Priority: redmine.Priority{Id: i + 1, Name: "High" + strconv.Itoa(i + 1)},
				Author:      redmine.Author{Id: i + 1, Name: "testUser" + strconv.Itoa(i + 1)},
				AssignedTo:  redmine.AssignedTo{Id: i + 1, Name: "testUser" + strconv.Itoa(i + 1)},
				Subject:     "test" + strconv.Itoa(i),
				Description: "testtesttesttest" + strconv.Itoa(i),
				StartDate:   "2020-01-01T00:00:00Z",
				CustomFields: redmine.CustomFields{&redmine.CustomField{
					Id:       i + 1,
					Name:     "customField" + strconv.Itoa(i + 1),
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
		}
		issueJsons10000[i] = &ij
	}
	_, _ = redmine.ListStatusId(issueJsons10000, "")
}

func BenchmarkListPriorityId(b *testing.B) {
	issueJsons10000 := make(redmine.Issues, 10000)
	for i := 0; i <10000; i++ {
		var ij redmine.Issue
		if i % 2 == 0 {
			ij = redmine.Issue{
				Id:      i,
				Project: redmine.Project{Id: i, Name: "testProject" + strconv.Itoa(i)},
				Tracker: redmine.Tracker{Id: i, Name: "doing" + strconv.Itoa(i)},
				Status:  redmine.Status{Id: i, Name: "test" + strconv.Itoa(i)}, Priority: redmine.Priority{Id: i, Name: "High" + strconv.Itoa(i)},
				Author:      redmine.Author{Id: i, Name: "testUser" + strconv.Itoa(i)},
				AssignedTo:  redmine.AssignedTo{Id: i, Name: "testUser" + strconv.Itoa(i)},
				Subject:     "test1" + strconv.Itoa(i),
				Description: "testtesttesttest" + strconv.Itoa(i),
				StartDate:   "2020-01-01T00:00:00Z",
				CustomFields: redmine.CustomFields{&redmine.CustomField{
					Id:       i,
					Name:     "customField" + strconv.Itoa(i),
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
		} else {
			ij = redmine.Issue{
				Id:      i,
				Project: redmine.Project{Id: i + 1, Name: "testProject" + strconv.Itoa(i + 1)},
				Tracker: redmine.Tracker{Id: i + 1, Name: "doing" + strconv.Itoa(i + 1)},
				Status:  redmine.Status{Id: i + 1, Name: "test" + strconv.Itoa(i + 1)}, Priority: redmine.Priority{Id: i + 1, Name: "High" + strconv.Itoa(i + 1)},
				Author:      redmine.Author{Id: i + 1, Name: "testUser" + strconv.Itoa(i + 1)},
				AssignedTo:  redmine.AssignedTo{Id: i + 1, Name: "testUser" + strconv.Itoa(i + 1)},
				Subject:     "test" + strconv.Itoa(i),
				Description: "testtesttesttest" + strconv.Itoa(i),
				StartDate:   "2020-01-01T00:00:00Z",
				CustomFields: redmine.CustomFields{&redmine.CustomField{
					Id:       i + 1,
					Name:     "customField" + strconv.Itoa(i + 1),
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
		}
		issueJsons10000[i] = &ij
	}
	_, _ = redmine.ListPriorityId(issueJsons10000, "")
}

func BenchmarkListUserIdAssignedTo(b *testing.B) {
	issueJsons10000 := make(redmine.Issues, 10000)
	for i := 0; i <10000; i++ {
		var ij redmine.Issue
		if i % 2 == 0 {
			ij = redmine.Issue{
				Id:      i,
				Project: redmine.Project{Id: i, Name: "testProject" + strconv.Itoa(i)},
				Tracker: redmine.Tracker{Id: i, Name: "doing" + strconv.Itoa(i)},
				Status:  redmine.Status{Id: i, Name: "test" + strconv.Itoa(i)}, Priority: redmine.Priority{Id: i, Name: "High" + strconv.Itoa(i)},
				Author:      redmine.Author{Id: i, Name: "testUser" + strconv.Itoa(i)},
				AssignedTo:  redmine.AssignedTo{Id: i, Name: "testUser" + strconv.Itoa(i)},
				Subject:     "test1" + strconv.Itoa(i),
				Description: "testtesttesttest" + strconv.Itoa(i),
				StartDate:   "2020-01-01T00:00:00Z",
				CustomFields: redmine.CustomFields{&redmine.CustomField{
					Id:       i,
					Name:     "customField" + strconv.Itoa(i),
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
		} else {
			ij = redmine.Issue{
				Id:      i,
				Project: redmine.Project{Id: i + 1, Name: "testProject" + strconv.Itoa(i + 1)},
				Tracker: redmine.Tracker{Id: i + 1, Name: "doing" + strconv.Itoa(i + 1)},
				Status:  redmine.Status{Id: i + 1, Name: "test" + strconv.Itoa(i + 1)}, Priority: redmine.Priority{Id: i + 1, Name: "High" + strconv.Itoa(i + 1)},
				Author:      redmine.Author{Id: i + 1, Name: "testUser" + strconv.Itoa(i + 1)},
				AssignedTo:  redmine.AssignedTo{Id: i + 1, Name: "testUser" + strconv.Itoa(i + 1)},
				Subject:     "test" + strconv.Itoa(i),
				Description: "testtesttesttest" + strconv.Itoa(i),
				StartDate:   "2020-01-01T00:00:00Z",
				CustomFields: redmine.CustomFields{&redmine.CustomField{
					Id:       i + 1,
					Name:     "customField" + strconv.Itoa(i + 1),
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
		}
		issueJsons10000[i] = &ij
	}
	_, _ = redmine.ListUserIdAssignedTo(issueJsons10000, "")
}

func BenchmarkListCustomFieldsId(b *testing.B) {
	issueJsons10000 := make(redmine.Issues, 10000)
	for i := 0; i <10000; i++ {
		var ij redmine.Issue
		if i % 2 == 0 {
			ij = redmine.Issue{
				Id:      i,
				Project: redmine.Project{Id: i, Name: "testProject" + strconv.Itoa(i)},
				Tracker: redmine.Tracker{Id: i, Name: "doing" + strconv.Itoa(i)},
				Status:  redmine.Status{Id: i, Name: "test" + strconv.Itoa(i)}, Priority: redmine.Priority{Id: i, Name: "High" + strconv.Itoa(i)},
				Author:      redmine.Author{Id: i, Name: "testUser" + strconv.Itoa(i)},
				AssignedTo:  redmine.AssignedTo{Id: i, Name: "testUser" + strconv.Itoa(i)},
				Subject:     "test1" + strconv.Itoa(i),
				Description: "testtesttesttest" + strconv.Itoa(i),
				StartDate:   "2020-01-01T00:00:00Z",
				CustomFields: redmine.CustomFields{&redmine.CustomField{
					Id:       i,
					Name:     "customField" + strconv.Itoa(i),
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
		} else {
			ij = redmine.Issue{
				Id:      i,
				Project: redmine.Project{Id: i + 1, Name: "testProject" + strconv.Itoa(i + 1)},
				Tracker: redmine.Tracker{Id: i + 1, Name: "doing" + strconv.Itoa(i + 1)},
				Status:  redmine.Status{Id: i + 1, Name: "test" + strconv.Itoa(i + 1)}, Priority: redmine.Priority{Id: i + 1, Name: "High" + strconv.Itoa(i + 1)},
				Author:      redmine.Author{Id: i + 1, Name: "testUser" + strconv.Itoa(i + 1)},
				AssignedTo:  redmine.AssignedTo{Id: i + 1, Name: "testUser" + strconv.Itoa(i + 1)},
				Subject:     "test" + strconv.Itoa(i),
				Description: "testtesttesttest" + strconv.Itoa(i),
				StartDate:   "2020-01-01T00:00:00Z",
				CustomFields: redmine.CustomFields{&redmine.CustomField{
					Id:       i + 1,
					Name:     "customField" + strconv.Itoa(i + 1),
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
		}
		issueJsons10000[i] = &ij
	}
	_, _ = redmine.ListCustomFieldsId(issueJsons10000, "")
}
