package redmine

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/utils"
	"strconv"
)

var dat map[string]interface{}

type Project struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Tracker struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Status struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Priority struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Author struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type AssignedTo struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type CustomField struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Multiple bool `json:"multiple"`
	Value interface{} `json:"value"`
}

type CustomFields []*CustomField

type Attachment struct {
	Id int `json:"id"`
	FileName string `json:"filename"`
	FileSize int64 `json:"filesize"`
	Description string `json:"description"`
	ContentUrl string `json:"content_url"`
	Author Author
	CreatedOn string `json:"created_on"`
}

type Attachments []*Attachment

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Detail struct {
	Property string `json:"property"`
	Name string `json:"name"`
	NewValue string `json:"new_value"`
	OldValue string `json:"old_value"`
}

type Details []*Detail

type Journal struct {
	Id int `json:"id"`
	User User `json:"user"`
	Notes string `json:"notes"`
	CreatedOn string `json:"created_on"`
	Details Details `json:"details"`
}

type Journals []*Journal

type Watcher struct {
	Id int
	Name string
}

type Watchers []*Watcher

type Issue struct{
	Id int `json:"id"`
	Project Project `json:"project"`
	Tracker Tracker `json:"tracker"`
	Status Status `json:"status"`
	Priority Priority `json:"priority"`
	Author Author `json:"author"`
	AssignedTo AssignedTo `json:"assigned_to"`
	Subject string `json:"subject"`
	Description string `json:"description"`
	StartDate string `json:"start_date"`
	DueDate string `json:"due_date"`
	DoneRatio int `json:"done_ratio"`
	CustomFields CustomFields `json:"custom_fields"`
	IsPrivate bool `json:"is_private"`
	EstimatedHours string `json:"estimated_hours"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
	ClosedOn string `json:"closed_on"`
	Attachments Attachments `json:"attachments"`
	Journals Journals `json:"journals"`
	Watchers Watchers `json:"watchers"`
}

type Issues []*Issue

type Uploads struct {
	Token string `json:"token,omitempty"`
	FileName string `json:"filename,omitempty"`
	ContentType string `json:"content_type,omitempty"`
}

type IssueParam struct {
	ProjectId     int          `json:"project_id,omitempty"`
	TrackerId     int          `json:"tracker_id,omitempty"`
	StatusId      int          `json:"status_id,omitempty"`
	PriorityId    int          `json:"priority_id,omitempty"`
	Subject       string       `json:"subject,omitempty"`
	Description   string       `json:"description,omitempty"`
	AssignedToId  int          `json:"assigned_to_id,omitempty"`
	ParentIssueId int          `json:"parent_issue_id,omitempty"`
	CustomFields  CustomFields `json:"custom_fields,omitempty"`
	Notes string `json:"notes,omitempty"`
	Uploads []Uploads `json:"uploads,omitempty"`
}

type IssueParamJson struct {
	Issue IssueParam `json:"issue"`
}

type FileParam struct {
	FileName string
	ContentType string
	Contents []byte
	Token string
}

var issuesResult struct {
	Issues Issues `json:"issues"`
	TotalCount int `json:"total_count"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
}

var issueResult struct {
	Issue Issue `json:"issue"`
}

var UploadResult struct {
	Upload struct {
		Token string `json:"token"`
	}
}

func unmarshalByteIssue(content []byte) (Issue, error) {
	var emptyIssue Issue
	err := json.Unmarshal(content, &issuesResult)
	if err != nil {
		return emptyIssue, err
	}
	return issueResult.Issue, nil
}

func GetIssues(url, key string, projectId int) (Issues, error){
	var issuesUrl string
	if projectId == 0 {
		issuesUrl = url + "/issues.json?key=" + key + "&limit=1&offset=0&status_id=*"
	} else {
		issuesUrl = url + "/issues.json?key=" + key + "&limit=1&offset=0&status_id=*&project_id=" + strconv.Itoa(projectId)
	}
	var issues Issues
	body, err := utils.Get(issuesUrl)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &issuesResult)
	if err != nil {
		return nil, err
	}
	fmt.Println(issuesResult.TotalCount)
	for offset := 0; offset < issuesResult.TotalCount; offset+=100 {
		if projectId == 0 {
			issuesUrl = url + "/issues.json?key=" + key + "&limit=100&offset=" + strconv.Itoa(offset) + "&status_id=*&sort=updated_on:asc"
		} else {
			issuesUrl = url + "/issues.json?key=" + key + "&limit=100&offset=" + strconv.Itoa(offset) + "&status_id=*&sort=updated_on:asc&project_id=" + strconv.Itoa(projectId)
		}
		body, err := utils.Get(issuesUrl)
		if err != nil {
			return nil, err
		}
		err2 := json.Unmarshal(body, &issuesResult)
		if err2 != nil {
			panic(err2)
		}
		issues = append(issues, issuesResult.Issues...)
	}
	return issues, nil
}

func GetIssue(url, key string, id int) (Issue, error) {
	var issue Issue
	body, err := utils.Get(url + "/issues/" + strconv.Itoa(id) + ".json?key=" + key + "&include=attachments,journals,watchers")
	if err != nil {
		return issue, err
	}
	err = json.Unmarshal(body, &issueResult)
	if err != nil {
		return issue, err
	}
	return issueResult.Issue, nil
}

func DownloadAttachmentFiles(key string, attachments Attachments) ([][]byte, error){
	var result [][]byte
	for _, file := range attachments {
		body, err := utils.Get(file.ContentUrl + "?key=" + key)
		if err != nil {
			return nil, err
		}
		result = append(result, body)
	}
	return result, nil
}

func CreateIssueFromByteSlice(content []byte) (*Issue, error) {
	err := json.Unmarshal(content, &issueResult)
	if err != nil {
		return nil, err
	}
	return &issueResult.Issue, nil
}

func CreateIssueParam(issue Issue, uploadFiles []FileParam) IssueParam {
	var issueParam IssueParam
	if issue.Attachments != nil {
		var uploads []Uploads
		for _, v := range uploadFiles {
			uploads = append(uploads, Uploads{FileName: v.FileName, ContentType: v.ContentType, Token: v.Token})
		}
		issueParam = IssueParam{
			ProjectId: issue.Project.Id,
			TrackerId: issue.Tracker.Id,
			StatusId: issue.Status.Id,
			PriorityId: issue.Priority.Id,
			AssignedToId: issue.AssignedTo.Id,
			Subject: issue.Subject,
			Description: issue.Description,
			CustomFields: issue.CustomFields,
			Uploads: uploads}
	} else {
		issueParam = IssueParam{
			ProjectId: issue.Project.Id,
			TrackerId: issue.Tracker.Id,
			StatusId: issue.Status.Id,
			PriorityId: issue.Priority.Id,
			AssignedToId: issue.AssignedTo.Id,
			Subject: issue.Subject,
			Description: issue.Description,
			CustomFields: issue.CustomFields}
	}
	return issueParam
}

func CreateIssue(url, key string, issue IssueParam) (int, error) {
	issueJson, err := json.Marshal(IssueParamJson{Issue: issue})
	if err != nil {
		return 0, err
	}
	fmt.Println(string(issueJson))
	body, err := utils.Post(url + "/issues.json?key=" + key, "application/json", issueJson)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(body, &issueResult)
	if err != nil {
		return 0, err
	}
	return issueResult.Issue.Id, nil
}

func UpdateIssueJournals(url, key string, id int, journals []string) error {
	for _, journal := range journals {
		issue := IssueParam{Notes: journal}
		issueJson, err := json.Marshal(IssueParamJson{Issue: issue})
		if err != nil {
			return err
		}
		fmt.Println(string(issueJson))
		err = utils.Put(url + "/issues/" + strconv.Itoa(id) + ".json?key=" + key, "application/json", issueJson)
		if err != nil {
			return err
		}
	}
	return nil
}

func UploadAttachmentFiles(url, key string, files []FileParam) ([]FileParam, error) {
	var newFiles []FileParam
	for _, file := range files {
		body, err := utils.Post(url + "/uploads.json?key=" + key + "&filename=" + file.FileName, "application/octet-stream", file.Contents)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, &UploadResult)
		if err != nil {
			return nil, err
		}
		newFiles = append(newFiles, FileParam{FileName: file.FileName, ContentType: file.ContentType, Contents: file.Contents, Token: UploadResult.Upload.Token})
	}
	return newFiles, nil
}

func CreateJournalStrings(issue Issue) []string {
	var notes []string
	for _, journal := range issue.Journals {
		notes = append(notes, journal.Notes)
	}
	return notes
}
