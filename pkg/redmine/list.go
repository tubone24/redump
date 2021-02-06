package redmine

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/tubone24/redump/pkg/utils"
	"github.com/goccy/go-json"
)

func ListProjectId(issues Issues, filename string) ([]Project, error) {
	var result []Project
	var unMarshalProject Project
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
	if filename == "" {
		filename = "data/project_id.json"
	}
	err = utils.WriteFile(filename, projectJson)
	return result, nil
}

func ListTrackerId(issues Issues, filename string) ([]Tracker, error) {
	s2 := mapset.NewSet()
	var result []Tracker
	var unMarshalTracker Tracker
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
	if filename == "" {
		filename = "data/tracker_id.json"
	}
	err = utils.WriteFile(filename, trackerJson)
	return result, nil
}

func ListStatusId(issues Issues, filename string) ([]Status, error) {
	s2 := mapset.NewSet()
	var result []Status
	var unMarshalStatus Status
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
	if filename == "" {
		filename = "data/status_id.json"
	}
	err = utils.WriteFile(filename, statusJson)
	return result, nil
}

func ListPriorityId(issues Issues, filename string) ([]Priority, error) {
	s2 := mapset.NewSet()
	var result []Priority
	var unMarshalPriority Priority
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
	if filename == "" {
		filename = "data/priority_id.json"
	}
	err = utils.WriteFile(filename, priorityJson)
	return result, nil
}

func ListUserIdAssignedTo(issues Issues, filename string) ([]AssignedTo, error) {
	s2 := mapset.NewSet()
	var result []AssignedTo
	var unMarshalAssignedTo AssignedTo
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
	if filename == "" {
		filename = "data/userId.json"
	}
	err = utils.WriteFile(filename, userJson)
	return result, nil
}

func ListCustomFieldsId(issues Issues, filename string) ([]CustomField, error) {
	s2 := mapset.NewSet()
	var result []CustomField
	var unMarshalCustomField CustomField
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
	if filename == "" {
		filename = "data/custom_fields_id.json"
	}
	err = utils.WriteFile(filename, customFieldsJson)
	return result, nil
}
