package cmd

import (
	"github.com/tubone24/redump/pkg/config"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/tubone24/redump/pkg/utils"
	"github.com/goccy/go-json"
	"github.com/deckarep/golang-set"
)

func ListAll(projectId int) error {
	cfg, err := config.GetConfig("")
	if err != nil {
		return err
	}
	issues, err := redmine.GetIssues(cfg.ServerConfig.Url, cfg.ServerConfig.Key, projectId, cfg.ServerConfig.Timeout, nil)
	_, err = ListProjectId(issues)
	if err != nil {
		return err
	}
	_, err = ListTrackerId(issues)
	if err != nil {
		return err
	}
	_, err = ListStatusId(issues)
	if err != nil {
		return err
	}
	_, err = ListPriorityId(issues)
	if err != nil {
		return err
	}
	_, err = ListUserIdAssignedTo(issues)
	if err != nil {
		return err
	}
	_, err = ListCustomFieldsId(issues)
	if err != nil {
		return err
	}
	return nil
}

func ListProjectId(issues redmine.Issues) ([]redmine.Project, error) {
	var result []redmine.Project
	var unMarshalProject redmine.Project
	s2 := mapset.NewSet()
	for _, issue := range issues {
		// result = append(result, issue.Project)
		projectIdJson, err := json.Marshal(issue.Project)
		if err != nil {
			return nil, err
		}
		s2.Add(string(projectIdJson))
	}
	for _, v := range s2.ToSlice() {
		err := json.Unmarshal([]byte(v.(string)), &unMarshalProject)
		if err != nil {
			return nil, err
		}
		result = append(result, unMarshalProject)
	}
	projectJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = utils.WriteFile("data/project_id.json", projectJson)
	return result, nil
}

func ListTrackerId(issues redmine.Issues) ([]redmine.Tracker, error) {
	s2 := mapset.NewSet()
	var result []redmine.Tracker
	var unMarshalTracker redmine.Tracker
	for _, issue := range issues {
		trackerJson, err := json.Marshal(issue.Tracker)
		if err != nil {
			return nil, err
		}
		s2.Add(string(trackerJson))
	}
	for _, v := range s2.ToSlice() {
		err := json.Unmarshal([]byte(v.(string)), &unMarshalTracker)
		if err != nil {
			return nil, err
		}
		result = append(result, unMarshalTracker)
	}
	trackerJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = utils.WriteFile("data/tracker_id.json", trackerJson)
	return result, nil
}

func ListStatusId(issues redmine.Issues) ([]redmine.Status, error) {
	s2 := mapset.NewSet()
	var result []redmine.Status
	var unMarshalStatus redmine.Status
	for _, issue := range issues {
		statusJson, err := json.Marshal(issue.Status)
		if err != nil {
			return nil, err
		}
		s2.Add(string(statusJson))
	}
	for _, v := range s2.ToSlice() {
		err := json.Unmarshal([]byte(v.(string)), &unMarshalStatus)
		if err != nil {
			return nil, err
		}
		result = append(result, unMarshalStatus)
	}
	statusJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = utils.WriteFile("data/status_id.json", statusJson)
	return result, nil
}

func ListPriorityId(issues redmine.Issues) ([]redmine.Priority, error) {
	s2 := mapset.NewSet()
	var result []redmine.Priority
	var unMarshalPriority redmine.Priority
	for _, issue := range issues {
		priorityJson, err := json.Marshal(issue.Priority)
		if err != nil {
			return nil, err
		}
		s2.Add(string(priorityJson))
	}
	for _, v := range s2.ToSlice() {
		err := json.Unmarshal([]byte(v.(string)), &unMarshalPriority)
		if err != nil {
			return nil, err
		}
		result = append(result, unMarshalPriority)
	}
	priorityJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = utils.WriteFile("data/priority_id.json", priorityJson)
	return result, nil
}

func ListUserIdAssignedTo(issues redmine.Issues) ([]redmine.AssignedTo, error) {
	s2 := mapset.NewSet()
	var result []redmine.AssignedTo
	var unMarshalAssignedTo redmine.AssignedTo
	for _, issue := range issues {
		assignedToJson, err := json.Marshal(issue.AssignedTo)
		if err != nil {
			return nil, err
		}
		s2.Add(string(assignedToJson))
	}
	for _, v := range s2.ToSlice() {
		err := json.Unmarshal([]byte(v.(string)), &unMarshalAssignedTo)
		if err != nil {
			return nil, err
		}
		result = append(result, unMarshalAssignedTo)
	}
	userJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = utils.WriteFile("data/userId.json", userJson)
	return result, nil
}

func ListCustomFieldsId(issues redmine.Issues) ([]redmine.CustomField, error) {
	s2 := mapset.NewSet()
	var result []redmine.CustomField
	var unMarshalCustomField redmine.CustomField
	for _, issue := range issues {
		for _, customField := range issue.CustomFields {
			customFieldJson, err := json.Marshal(customField)
			if err != nil {
				return nil, err
			}
			s2.Add(string(customFieldJson))
		}
	}
	for _, v := range s2.ToSlice() {
		err := json.Unmarshal([]byte(v.(string)), &unMarshalCustomField)
		if err != nil {
			return nil, err
		}
		result = append(result, unMarshalCustomField)
	}
	customFieldsJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = utils.WriteFile("data/custom_fields_id.json", customFieldsJson)
	return result, nil
}
