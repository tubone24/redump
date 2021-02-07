package redmine

import (
	"github.com/tubone24/redump/pkg/config"
	"github.com/tubone24/redump/pkg/utils"
)

// ConvertNewEnv converts the information in the source Redmine ticket
// to that of the new Redmine ticket using the mappings that exist in Config.
// When silent mode is enabled, only the user specified in default will be assigned.
func ConvertNewEnv(issue Issue, conf config.Config, silent bool) (*Issue, error) {
	var newIssue Issue
	err := utils.DeepCopy(&newIssue, &issue)
	if err != nil {
		return nil, err
	}
	for _, mapping := range conf.Mappings {
		switch mapping.Name {
		case "project_id":
			newIssue.Project.Id = convertProjectId(&issue, mapping.Values, mapping.Default)
		case "tracker_id":
			newIssue.Tracker.Id = convertTrackerId(&issue, mapping.Values, mapping.Default)
		case "status_id":
			newIssue.Status.Id = convertStatusId(&issue, mapping.Values, mapping.Default)
		case "priority_id":
			newIssue.Priority.Id = convertPriorityId(&issue, mapping.Values, mapping.Default)
		case "user_id":
			newIssue.AssignedTo.Id = convertUserIdToAssignedTo(&issue, mapping.Values, mapping.Default, silent)
			for i, v := range convertWatcherId(&issue, mapping.Values, mapping.Default) {
				newIssue.Watchers[i].Id = v
			}
		case "custom_field_id":
			for i, v := range convertCustomFieldsId(&issue, mapping.Values, mapping.Default) {
				newIssue.CustomFields[i].Id = v
			}
		}
	}
	return &newIssue, nil
}

func convertProjectId(issue *Issue, conf []config.MappingValue, defaultValue int) int {
	for _, v := range conf {
		if issue.Project.Id == v.Old {
			return v.New
		}
	}
	return defaultValue
}

func convertTrackerId(issue *Issue, conf []config.MappingValue, defaultValue int) int {
	for _, v := range conf {
		if issue.Tracker.Id == v.Old {
			return v.New
		}
	}
	return defaultValue
}

func convertStatusId(issue *Issue, conf []config.MappingValue, defaultValue int) int {
	for _, v := range conf {
		if issue.Status.Id == v.Old {
			return v.New
		}
	}
	return defaultValue
}

func convertPriorityId(issue *Issue, conf []config.MappingValue, defaultValue int) int {
	for _, v := range conf {
		if issue.Priority.Id == v.Old {
			return v.New
		}
	}
	return defaultValue
}

func convertUserIdToAssignedTo(issue *Issue, conf []config.MappingValue, defaultValue int, silent bool) int {
	if silent {
		return defaultValue
	}
	for _, v := range conf {
		if issue.AssignedTo.Id == v.Old {
			return v.New
		}
	}
	return defaultValue
}

func convertCustomFieldsId(issue *Issue, conf []config.MappingValue, defaultValue int) []int {
	// var result []int
	result := make([]int, len(issue.CustomFields))
	if issue.CustomFields != nil {
		for i, v := range issue.CustomFields {
			for i2, v2 := range conf {
				if v.Id == v2.Old {
					//result = append(result, v2.New)
					result[i] = v2.New
					break
				} else {
					if i2+1 == len(conf) {
						//result = append(result, defaultValue)
						result[i] = defaultValue
					}
				}
			}
		}
	}
	return result
}

func convertWatcherId(issue *Issue, conf []config.MappingValue, defaultValue int) []int {
	// var result []int
	result := make([]int, len(issue.Watchers))
	if issue.Watchers != nil {
		for i, v := range issue.Watchers {
			for i2, v2 := range conf {
				if v.Id == v2.Old {
					//result = append(result, v2.New)
					result[i] = v2.New
					break
				} else {
					if i2+1 == len(conf) {
						//result = append(result, defaultValue)
						result[i] = defaultValue
					}
				}
			}
		}
	}
	return result
}
