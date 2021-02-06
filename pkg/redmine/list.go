package redmine

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/utils"
)

func ListProjectId(issues Issues, filename string) ([]Project, error) {
	var unMarshalProject Project
	s2 := mapset.NewSet()
	for _, issue := range issues {
		projectIdJson, err := json.Marshal(issue.Project)
		if err != nil {
			return nil, err
		}
		s2.Add(string(projectIdJson))
	}
	//var result []Project
	result := make([]Project, len(s2.ToSlice()))
	for i, v := range s2.ToSlice() {
		err := json.Unmarshal([]byte(v.(string)), &unMarshalProject)
		if err != nil {
			return nil, err
		}
		// result = append(result, unMarshalProject)
		result[i] = unMarshalProject
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
	var unMarshalTracker Tracker
	for _, issue := range issues {
		trackerJson, err := json.Marshal(issue.Tracker)
		if err != nil {
			return nil, err
		}
		s2.Add(string(trackerJson))
	}
	// var result []Tracker
	result := make([]Tracker, len(s2.ToSlice()))
	for i, v := range s2.ToSlice() {
		err := json.Unmarshal([]byte(v.(string)), &unMarshalTracker)
		if err != nil {
			return nil, err
		}
		// result = append(result, unMarshalTracker)
		result[i] = unMarshalTracker
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
	var unMarshalStatus Status
	for _, issue := range issues {
		statusJson, err := json.Marshal(issue.Status)
		if err != nil {
			return nil, err
		}
		s2.Add(string(statusJson))
	}
	// var result []Status
	result := make([]Status, len(s2.ToSlice()))
	for i, v := range s2.ToSlice() {
		err := json.Unmarshal([]byte(v.(string)), &unMarshalStatus)
		if err != nil {
			return nil, err
		}
		// result = append(result, unMarshalStatus)
		result[i] = unMarshalStatus
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
	var unMarshalPriority Priority
	for _, issue := range issues {
		priorityJson, err := json.Marshal(issue.Priority)
		if err != nil {
			return nil, err
		}
		s2.Add(string(priorityJson))
	}
	//var result []Priority
	result := make([]Priority, len(s2.ToSlice()))
	for i, v := range s2.ToSlice() {
		err := json.Unmarshal([]byte(v.(string)), &unMarshalPriority)
		if err != nil {
			return nil, err
		}
		// result = append(result, unMarshalPriority)
		result[i] = unMarshalPriority
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
	var unMarshalAssignedTo AssignedTo
	for _, issue := range issues {
		assignedToJson, err := json.Marshal(issue.AssignedTo)
		if err != nil {
			return nil, err
		}
		s2.Add(string(assignedToJson))
	}
	// var result []AssignedTo
	result := make([]AssignedTo, len(s2.ToSlice()))
	for i, v := range s2.ToSlice() {
		err := json.Unmarshal([]byte(v.(string)), &unMarshalAssignedTo)
		if err != nil {
			return nil, err
		}
		// result = append(result, unMarshalAssignedTo)
		result[i] = unMarshalAssignedTo
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
	//var result []CustomField
	result := make([]CustomField, len(s2.ToSlice()))
	for i, v := range s2.ToSlice() {
		err := json.Unmarshal([]byte(v.(string)), &unMarshalCustomField)
		if err != nil {
			return nil, err
		}
		//result = append(result, unMarshalCustomField)
		result[i] = unMarshalCustomField
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
