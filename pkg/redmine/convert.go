package redmine

import "github.com/tubone24/redump/pkg/config"

func ConvertNewEnv(issue Issue) (*Issue, error){
	conf, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	for _, mapping := range conf.Mappings {
		switch mapping.Name {
			case "project_id":
				convertProjectId(&issue, mapping.Values)
			case "tracker_id":
				convertTrackerId(&issue, mapping.Values, mapping.Default)
			case "status_id":
				convertStatusId(&issue, mapping.Values, mapping.Default)
			case "priority_id":
				convertPriorityId(&issue, mapping.Values, mapping.Default)
			case "user_id":
				convertUserIdToAssignedTo(&issue, mapping.Values, mapping.Default)
				// Watcherはそのうちやる
			case "custom_field_id":
				convertCustomFieldsId(&issue, mapping.Values)
		}
	}
	return &issue, nil
}

func convertProjectId (issue *Issue, conf []config.MappingValue) {
	for _, v := range conf {
		if issue.Project.Id == v.Old {
			issue.Project.Id = v.New
			break
		}
	}
}

func convertTrackerId (issue *Issue, conf []config.MappingValue, defaultValue int) {
	for _, v := range conf {
		if issue.Tracker.Id == v.Old {
			issue.Tracker.Id = v.New
			return
		}
	}
	issue.Tracker.Id = defaultValue
}

func convertStatusId (issue *Issue, conf []config.MappingValue, defaultValue int) {
	for _, v := range conf {
		if issue.Status.Id == v.Old {
			issue.Status.Id = v.New
			return
		}
	}
	issue.Status.Id = defaultValue
}

func convertPriorityId (issue *Issue, conf []config.MappingValue, defaultValue int) {
	for _, v := range conf {
		if issue.Priority.Id == v.Old {
			issue.Priority.Id = v.New
			return
		}
	}
	issue.Priority.Id = defaultValue
}

func convertUserIdToAssignedTo (issue *Issue, conf []config.MappingValue, defaultValue int) {
	for _, v := range conf {
		if issue.AssignedTo.Id == v.Old {
			issue.AssignedTo.Id = v.New
			return
		}
	}
	issue.AssignedTo.Id = defaultValue
}

func convertCustomFieldsId (issue *Issue, conf []config.MappingValue) {
	if issue.CustomFields != nil {
		for _, v := range conf {
			for _, v2 := range issue.CustomFields {
				if v2.Id == v.Old {
					v2.Id = v.New
					break
				}
			}
		}
	}
}
