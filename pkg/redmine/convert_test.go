package redmine_test

import (
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
	actual, err := redmine.ConvertNewEnv(issueJson, cfg)
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
	actual, err := redmine.ConvertNewEnv(issueJson, cfg)
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

func TestDeepCopy(t *testing.T) {
	var newIssueJson redmine.Issue
	err := redmine.DeepCopy(&newIssueJson, &issueJson)
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
	newIssueJson.Id = 50
	if issueJson.Id == newIssueJson.Id {
		t.Error("Destructive changes")
	}
}
