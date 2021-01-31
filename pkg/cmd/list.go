package cmd

import (
	"github.com/tubone24/redump/pkg/config"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/tubone24/redump/pkg/utils"
	"github.com/goccy/go-json"
)

func ListAll(projectId int) {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	issues, err := redmine.GetIssues(config.ServerConfig.Url, config.ServerConfig.Key, projectId)
	_, err = ListProjectId(issues)
	if err != nil {
		panic(err)
	}
	_, err = ListTrackerId(issues)
	if err != nil {
		panic(err)
	}
	_, err = ListStatusId(issues)
	if err != nil {
		panic(err)
	}
	_, err = ListPriorityId(issues)
	if err != nil {
		panic(err)
	}
	_, err = ListUserIdAssignedTo(issues)
	if err != nil {
		panic(err)
	}
	_, err = ListCustomFieldsId(issues)
	if err != nil {
		panic(err)
	}
}

func ListProjectId(issues redmine.Issues) ([]redmine.Project, error) {
	var result []redmine.Project
	for _, issue := range issues {
		result = append(result, issue.Project)
	}
	projectIdJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = utils.WriteFile("data/project_id.json", projectIdJson)
	return result, nil
}

func ListTrackerId(issues redmine.Issues) ([]redmine.Tracker, error) {
	var result []redmine.Tracker
	for _, issue := range issues {
		result = append(result, issue.Tracker)
	}
	trackerIdJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = utils.WriteFile("data/tracker_id.json", trackerIdJson)
	return result, nil
}

func ListStatusId(issues redmine.Issues) ([]redmine.Status, error) {
	var result []redmine.Status
	for _, issue := range issues {
		result = append(result, issue.Status)
	}
	statusIdJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = utils.WriteFile("data/status_id.json", statusIdJson)
	return result, nil
}

func ListPriorityId(issues redmine.Issues) ([]redmine.Priority, error) {
	var result []redmine.Priority
	for _, issue := range issues {
		result = append(result, issue.Priority)
	}
	priorityIdJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = utils.WriteFile("data/priority_id.json", priorityIdJson)
	return result, nil
}

func ListUserIdAssignedTo(issues redmine.Issues) ([]redmine.AssignedTo, error) {
	var result []redmine.AssignedTo
	for _, issue := range issues {
		result = append(result, issue.AssignedTo)
	}
	userIdJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = utils.WriteFile("data/userId.json", userIdJson)
	return result, nil
}

func ListCustomFieldsId(issues redmine.Issues) ([]redmine.CustomField, error) {
	var result []redmine.CustomField
	for _, issue := range issues {
		for _, customField := range issue.CustomFields {
			result = append(result, *customField)
		}
	}
	customFieldsIdJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = utils.WriteFile("data/custom_fields_id.json", customFieldsIdJson)
	return result, nil
}
