package redmine_test

import (
	"github.com/tubone24/redump/pkg/redmine"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestListProjectId(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "project_id.json")
	resp, err := redmine.ListProjectId(issuesJson, filename)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if resp[0].Id != issueJson.Project.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Project.Id, resp[0].Id)
	}
	if resp[0].Name != issueJson.Project.Name {
		t.Errorf("expected: %s, actual %s", issueJson.Project.Name, resp[0].Name)
	}
}

func TestListTrackerId(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "tracker_id.json")
	resp, err := redmine.ListTrackerId(issuesJson, filename)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if resp[0].Id != issueJson.Tracker.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Tracker.Id, resp[0].Id)
	}
	if resp[0].Name != issueJson.Tracker.Name {
		t.Errorf("expected: %s, actual %s", issueJson.Tracker.Name, resp[0].Name)
	}
}

func TestListStatusId(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "status_id.json")
	resp, err := redmine.ListStatusId(issuesJson, filename)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if resp[0].Id != issueJson.Status.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Status.Id, resp[0].Id)
	}
	if resp[0].Name != issueJson.Status.Name {
		t.Errorf("expected: %s, actual %s", issueJson.Status.Name, resp[0].Name)
	}
}

func TestListPriorityId(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "priority_id.json")
	resp, err := redmine.ListPriorityId(issuesJson, filename)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if resp[0].Id != issueJson.Priority.Id {
		t.Errorf("expected: %d, actual %d", issueJson.Priority.Id, resp[0].Id)
	}
	if resp[0].Name != issueJson.Priority.Name {
		t.Errorf("expected: %s, actual %s", issueJson.Priority.Name, resp[0].Name)
	}
}

func TestListUserIdAssignedTo(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "user_id.json")
	resp, err := redmine.ListUserIdAssignedTo(issuesJson, filename)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if resp[0].Id != issueJson.AssignedTo.Id {
		t.Errorf("expected: %d, actual %d", issueJson.AssignedTo.Id, resp[0].Id)
	}
	if resp[0].Name != issueJson.AssignedTo.Name {
		t.Errorf("expected: %s, actual %s", issueJson.AssignedTo.Name, resp[0].Name)
	}
}

func TestListCustomFieldsId(t *testing.T) {
	issuesJson := redmine.Issues{&issueJson}
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "custom_fields_id_list.json")
	resp, err := redmine.ListCustomFieldsId(issuesJson, filename)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if resp[0].Id != issueJson.CustomFields[0].Id {
		t.Errorf("expected: %d, actual %d", issueJson.CustomFields[0].Id, resp[0].Id)
	}
	if resp[0].Name != issueJson.CustomFields[0].Name {
		t.Errorf("expected: %s, actual %s", issueJson.CustomFields[0].Name, resp[0].Name)
	}
}
