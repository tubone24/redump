package redmine_test

import (
	"fmt"
	"github.com/tubone24/redump/pkg/config"
	"github.com/tubone24/redump/pkg/redmine"
	"testing"
)

func TestConvertNewEnv(t *testing.T) {
	cfg := config.Config{Mappings: []config.Mapping{
		{
			Name: "project_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "tracker_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "status_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "priority_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "user_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "custom_field_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
	}}
	actual, err := redmine.ConvertNewEnv(issueJson, cfg, false)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if actual.Project.Id != cfg.Mappings[0].Values[0].New {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[0].Values[0].New, actual.Project.Id)
	}
	if actual.Tracker.Id != cfg.Mappings[1].Values[0].New {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[1].Values[0].New, actual.Tracker.Id)
	}
	if actual.Status.Id != cfg.Mappings[2].Values[0].New {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[2].Values[0].New, actual.Status.Id)
	}
	if actual.Priority.Id != cfg.Mappings[3].Values[0].New {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[3].Values[0].New, actual.Priority.Id)
	}
	if actual.AssignedTo.Id != cfg.Mappings[4].Values[0].New {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[4].Values[0].New, actual.AssignedTo.Id)
	}
	if actual.CustomFields[0].Id != cfg.Mappings[5].Values[0].New {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[5].Values[0].New, actual.CustomFields[0].Id)
	}
}

func TestConvertNewEnvSilent(t *testing.T) {
	cfg := config.Config{Mappings: []config.Mapping{
		{
			Name: "project_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "tracker_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "status_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "priority_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "user_id",
			Default: 5,
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "custom_field_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
	}}
	actual, err := redmine.ConvertNewEnv(issueJson, cfg, true)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if actual.Project.Id != cfg.Mappings[0].Values[0].New {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[0].Values[0].New, actual.Project.Id)
	}
	if actual.Tracker.Id != cfg.Mappings[1].Values[0].New {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[1].Values[0].New, actual.Tracker.Id)
	}
	if actual.Status.Id != cfg.Mappings[2].Values[0].New {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[2].Values[0].New, actual.Status.Id)
	}
	if actual.Priority.Id != cfg.Mappings[3].Values[0].New {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[3].Values[0].New, actual.Priority.Id)
	}
	if actual.AssignedTo.Id != cfg.Mappings[4].Default {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[4].Default, actual.AssignedTo.Id)
	}
	if actual.CustomFields[0].Id != cfg.Mappings[5].Values[0].New {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[5].Values[0].New, actual.CustomFields[0].Id)
	}
}

func TestConvertNewEnvDefaultValue(t *testing.T) {
	cfg := config.Config{Mappings: []config.Mapping{
		{
			Name: "project_id",
			Default: 5,
			Values: []config.MappingValue{
				{
					Old: 100,
					New: 2,
				},
			},
		},
		{
			Name: "tracker_id",
			Default: 5,
			Values: []config.MappingValue{
				{
					Old: 100,
					New: 2,
				},
			},
		},
		{
			Name: "status_id",
			Default: 5,
			Values: []config.MappingValue{
				{
					Old: 100,
					New: 2,
				},
			},
		},
		{
			Name: "priority_id",
			Default: 5,
			Values: []config.MappingValue{
				{
					Old: 100,
					New: 2,
				},
			},
		},
		{
			Name: "user_id",
			Default: 5,
			Values: []config.MappingValue{
				{
					Old: 100,
					New: 2,
				},
			},
		},
		{
			Name: "custom_field_id",
			Default: 5,
			Values: []config.MappingValue{
				{
					Old: 100,
					New: 2,
				},
			},
		},
	}}
	actual, err := redmine.ConvertNewEnv(issueJson, cfg, false)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	if actual.Project.Id != cfg.Mappings[0].Default {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[0].Default, actual.Project.Id)
	}
	if actual.Tracker.Id != cfg.Mappings[1].Default {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[1].Default, actual.Tracker.Id)
	}
	if actual.Status.Id != cfg.Mappings[2].Default {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[2].Default, actual.Status.Id)
	}
	if actual.Priority.Id != cfg.Mappings[3].Default {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[3].Default, actual.Priority.Id)
	}
	if actual.AssignedTo.Id != cfg.Mappings[4].Default {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[4].Default, actual.AssignedTo.Id)
	}
	if actual.CustomFields[0].Id != cfg.Mappings[5].Default {
		t.Errorf("expected: %d, actual %d", cfg.Mappings[5].Default, actual.CustomFields[0].Id)
	}
}

func BenchmarkConvertNewEnv(b *testing.B) {
	cfg := config.Config{Mappings: []config.Mapping{
		{
			Name: "project_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "tracker_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "status_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "priority_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "user_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "custom_field_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
	}}
	b.ResetTimer()
	for i := 0; i < 100; i++ {
		_, _ = redmine.ConvertNewEnv(issueJson, cfg, false)
	}
}

func ExampleConvertNewEnv() {
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
					NewValue: "bbb"}},},
		},
		Watchers: redmine.Watchers{&redmine.Watcher{
			Id: 1, Name: "testUser"}, &redmine.Watcher{Id: 2, Name: "testUser2"}, &redmine.Watcher{Id: 3, Name: "testUser3"},},
	}
	cfg := config.Config{Mappings: []config.Mapping{
		{
			Name: "project_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "tracker_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "status_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "priority_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "user_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
		{
			Name: "custom_field_id",
			Values: []config.MappingValue{
				{
					Old: 1,
					New: 2,
				},
			},
		},
	}}
	resp, _ := redmine.ConvertNewEnv(exampleIssue, cfg, false)
	fmt.Println(exampleIssue.Project.Id)
	// Change to 1 to 2
	fmt.Println(resp.Project.Id)
}
