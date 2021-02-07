package utils_test

import (
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/tubone24/redump/pkg/utils"
	"testing"
)

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

func TestDeepCopy(t *testing.T) {
	var newIssueJson redmine.Issue
	err := utils.DeepCopy(&newIssueJson, &issueJson)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	newIssueJson.Id = 50
	if issueJson.Id == newIssueJson.Id {
		t.Error("Destructive changes")
	}
}

func TestDeepCopyInvalidSrc(t *testing.T) {
	var newIssueJson redmine.Issue
	err := utils.DeepCopy(&newIssueJson, "")
	if err == nil {
		t.Error("Error does not occurred")
	}
}

func TestDeepCopyInvalidDst(t *testing.T) {
	err := utils.DeepCopy("", &issueJson)
	if err == nil {
		t.Error("Error does not occurred")
	}
}

func ExampleDeepCopy() {
	source := redmine.Issue{
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
	var destination redmine.Issue
	// Pass a pointer to a structure with the same type
	_ = utils.DeepCopy(&destination, &source)
}
