// Redmine Package is a set of accessors that provide you to get, update, and delete Redmine Issues, Wikis, files, and project-wide configuration values.
package redmine

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/utils"
	"net/http"
	"net/url"
	"strconv"
)

type Project struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Tracker struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Status struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Priority struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type AssignedTo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CustomField struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Multiple bool        `json:"multiple"`
	Value    interface{} `json:"value"`
}

type CustomFields []*CustomField

type Attachment struct {
	Id          int    `json:"id"`
	FileName    string `json:"filename"`
	FileSize    int64  `json:"filesize"`
	Description string `json:"description"`
	ContentUrl  string `json:"content_url"`
	Author      Author
	CreatedOn   string `json:"created_on"`
}

type Attachments []*Attachment

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Detail struct {
	Property string `json:"property"`
	Name     string `json:"name"`
	NewValue string `json:"new_value"`
	OldValue string `json:"old_value"`
}

type Details []*Detail

type Journal struct {
	Id        int     `json:"id"`
	User      User    `json:"user"`
	Notes     string  `json:"notes"`
	CreatedOn string  `json:"created_on"`
	Details   Details `json:"details"`
}

type Journals []*Journal

type Watcher struct {
	Id   int
	Name string
}

type Watchers []*Watcher

type Issue struct {
	Id             int          `json:"id"`
	Project        Project      `json:"project"`
	Tracker        Tracker      `json:"tracker"`
	Status         Status       `json:"status"`
	Priority       Priority     `json:"priority"`
	Author         Author       `json:"author"`
	AssignedTo     AssignedTo   `json:"assigned_to"`
	Subject        string       `json:"subject"`
	Description    string       `json:"description"`
	StartDate      string       `json:"start_date"`
	DueDate        string       `json:"due_date"`
	DoneRatio      int          `json:"done_ratio"`
	CustomFields   CustomFields `json:"custom_fields"`
	IsPrivate      bool         `json:"is_private"`
	EstimatedHours string       `json:"estimated_hours"`
	CreatedOn      string       `json:"created_on"`
	UpdatedOn      string       `json:"updated_on"`
	ClosedOn       string       `json:"closed_on"`
	Attachments    Attachments  `json:"attachments"`
	Journals       Journals     `json:"journals"`
	Watchers       Watchers     `json:"watchers"`
}

type Issues []*Issue

type Uploads struct {
	Token       string `json:"token,omitempty"`
	FileName    string `json:"filename,omitempty"`
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
	Notes         string       `json:"notes,omitempty"`
	Uploads       []Uploads    `json:"uploads,omitempty"`
}

type IssueParamJson struct {
	Issue IssueParam `json:"issue"`
}

type FileParam struct {
	FileName    string
	ContentType string
	Contents    []byte
	Token       string
}

type WatcherParam struct {
	UserId int `json:"user_id"`
}

var issuesResult struct {
	Issues     Issues `json:"issues"`
	TotalCount int    `json:"total_count"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
}

var issueResult struct {
	Issue Issue `json:"issue"`
}

var UploadResult struct {
	Upload struct {
		Token string `json:"token"`
	}
}

// UnmarshalByteIssue is function to create the structure of an Issue by receiving a byte slice.
// If an invalid byte slice is received, an empty Issue structure will be returned with an error.
// Deprecated: This function is deprecated
func UnmarshalByteIssue(content []byte) (Issue, error) {
	var emptyIssue Issue
	err := json.Unmarshal(content, &issuesResult)
	if err != nil {
		return emptyIssue, err
	}
	return issueResult.Issue, nil
}

// GetIssues is function that you can get all issues of Redmine.
// However, you can't get detailed information such as Watchers and Journals.
// If you want to get them, you have to specify the Issue ID and get them individually for GetIssue.
func GetIssues(url, key string, projectId, timeout int, customClient *http.Client) (Issues, error) {
	var issuesUrl string
	if projectId == 0 {
		issuesUrl = url + "/issues.json?key=" + key + "&limit=1&offset=0&status_id=*"
	} else {
		issuesUrl = url + "/issues.json?key=" + key + "&limit=1&offset=0&status_id=*&project_id=" + strconv.Itoa(projectId)
	}
	var issues Issues
	var client *utils.Api
	if customClient == nil {
		client = utils.NewHttpClient(timeout)
	} else {
		client = utils.NewHttpClient(timeout, utils.OptionHTTPClient(customClient))
	}
	body, err := client.Get(issuesUrl)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &issuesResult)
	if err != nil {
		return nil, err
	}
	fmt.Println(issuesResult.TotalCount)
	for offset := 0; offset < issuesResult.TotalCount; offset += 100 {
		if projectId == 0 {
			issuesUrl = url + "/issues.json?key=" + key + "&limit=100&offset=" + strconv.Itoa(offset) + "&status_id=*&sort=updated_on:asc"
		} else {
			issuesUrl = url + "/issues.json?key=" + key + "&limit=100&offset=" + strconv.Itoa(offset) + "&status_id=*&sort=updated_on:asc&project_id=" + strconv.Itoa(projectId)
		}
		body, err := client.Get(issuesUrl)
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

// GetIssue is function that you can retrieve the details by specifying the Issue ID.
func GetIssue(url, key string, id, timeout int, customClient *http.Client) (Issue, error) {
	var issue Issue
	var client *utils.Api
	if customClient == nil {
		client = utils.NewHttpClient(timeout)
	} else {
		client = utils.NewHttpClient(timeout, utils.OptionHTTPClient(customClient))
	}
	body, err := client.Get(url + "/issues/" + strconv.Itoa(id) + ".json?key=" + key + "&include=attachments,journals,watchers")
	if err != nil {
		return issue, err
	}
	err = json.Unmarshal(body, &issueResult)
	if err != nil {
		return issue, err
	}
	return issueResult.Issue, nil
}

// DownloadAttachmentFiles is function that you pass the Attachment structure,
// you can store all the included attachments locally and return a byte slice of the attachment along with it
func DownloadAttachmentFiles(key string, timeout int, attachments Attachments, customClient *http.Client) ([][]byte, error) {
	//var result [][]byte
	result := make([][]byte, len(attachments))
	var client *utils.Api
	if customClient == nil {
		client = utils.NewHttpClient(timeout)
	} else {
		client = utils.NewHttpClient(timeout, utils.OptionHTTPClient(customClient))
	}
	for i, file := range attachments {
		body, err := client.Get(file.ContentUrl + "?key=" + key)
		if err != nil {
			return nil, err
		}
		//result = append(result, body)
		result[i] = body
	}
	return result, nil
}

// CreateIssueFromByteSlice is function to create the structure of an Issue by receiving a byte slice.
// If an invalid byte slice is received, an empty Issue structure will be returned with an error.
func CreateIssueFromByteSlice(content []byte) (*Issue, error) {
	var issue Issue
	err := json.Unmarshal(content, &issue)
	if err != nil {
		return nil, err
	}
	return &issue, nil
}

// CreateIssueParam is
func CreateIssueParam(issue Issue, uploadFiles []FileParam) IssueParam {
	var issueParam IssueParam
	if issue.Attachments != nil {
		//var uploads []Uploads
		uploads := make([]Uploads, len(uploadFiles))
		for i, v := range uploadFiles {
			// uploads = append(uploads, Uploads{FileName: v.FileName, ContentType: v.ContentType, Token: v.Token})
			uploads[i] = Uploads{FileName: v.FileName, ContentType: v.ContentType, Token: v.Token}
		}
		issueParam = IssueParam{
			ProjectId:    issue.Project.Id,
			TrackerId:    issue.Tracker.Id,
			StatusId:     issue.Status.Id,
			PriorityId:   issue.Priority.Id,
			AssignedToId: issue.AssignedTo.Id,
			Subject:      issue.Subject,
			Description:  issue.Description,
			CustomFields: issue.CustomFields,
			Uploads:      uploads}
	} else {
		issueParam = IssueParam{
			ProjectId:    issue.Project.Id,
			TrackerId:    issue.Tracker.Id,
			StatusId:     issue.Status.Id,
			PriorityId:   issue.Priority.Id,
			AssignedToId: issue.AssignedTo.Id,
			Subject:      issue.Subject,
			Description:  issue.Description,
			CustomFields: issue.CustomFields}
	}
	return issueParam
}

func CreateIssue(url, key string, timeout int, issue IssueParam, customClient *http.Client) (int, error) {
	issueJson, err := json.Marshal(IssueParamJson{Issue: issue})
	if err != nil {
		return 0, err
	}
	fmt.Println(string(issueJson))
	var client *utils.Api
	if customClient == nil {
		client = utils.NewHttpClient(timeout)
	} else {
		client = utils.NewHttpClient(timeout, utils.OptionHTTPClient(customClient))
	}
	body, err := client.Post(url+"/issues.json?key="+key, "application/json", issueJson)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(body, &issueResult)
	if err != nil {
		return 0, err
	}
	return issueResult.Issue.Id, nil
}

func UpdateIssueJournals(url, key string, id, timeout int, journals []string, customClient *http.Client) error {
	for _, journal := range journals {
		if journal == "" {
			continue
		}
		issue := IssueParam{Notes: journal}
		issueJson, err := json.Marshal(IssueParamJson{Issue: issue})
		if err != nil {
			return err
		}
		fmt.Println(string(issueJson))
		var client *utils.Api
		if customClient == nil {
			client = utils.NewHttpClient(timeout)
		} else {
			client = utils.NewHttpClient(timeout, utils.OptionHTTPClient(customClient))
		}
		err = client.Put(url+"/issues/"+strconv.Itoa(id)+".json?key="+key, "application/json", issueJson)
		if err != nil {
			return err
		}
	}
	return nil
}

func UploadAttachmentFiles(u, key string, timeout int, files []FileParam, customClient *http.Client) ([]FileParam, error) {
	// var newFiles []FileParam
	newFiles := make([]FileParam, len(files))
	for i, file := range files {
		params := url.Values{}
		params.Set("key", key)
		params.Add("filename", file.FileName)
		var client *utils.Api
		if customClient == nil {
			client = utils.NewHttpClient(timeout)
		} else {
			client = utils.NewHttpClient(timeout, utils.OptionHTTPClient(customClient))
		}
		body, err := client.Post(u+"/uploads.json?"+params.Encode(), "application/octet-stream", file.Contents)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, &UploadResult)
		if err != nil {
			return nil, err
		}
		// newFiles = append(newFiles, FileParam{FileName: file.FileName, ContentType: file.ContentType, Contents: file.Contents, Token: UploadResult.Upload.Token})
		newFiles[i] =  FileParam{FileName: file.FileName, ContentType: file.ContentType, Contents: file.Contents, Token: UploadResult.Upload.Token}
	}
	return newFiles, nil
}

func CreateJournalStrings(issue Issue) []string {
	// var notes []string
	notes := make([]string, len(issue.Journals))
	for i, journal := range issue.Journals {
		// notes = append(notes, journal.Notes)
		notes[i] = journal.Notes
	}
	return notes
}

func UpdateWatchers(url, key string, id, timeout int, issue Issue, customClient *http.Client) error {
	watcherSet := mapset.NewSet()
	for _, watcher := range issue.Watchers {
		watcherSet.Add(watcher.Id)
	}
	for _, watcherId := range watcherSet.ToSlice() {
		watcherJson, err := json.Marshal(WatcherParam{UserId: watcherId.(int)})
		fmt.Println(string(watcherJson))
		if err != nil {
			return err
		}
		var client *utils.Api
		if customClient == nil {
			client = utils.NewHttpClient(timeout)
		} else {
			client = utils.NewHttpClient(timeout, utils.OptionHTTPClient(customClient))
		}
		_, err = client.Post(url+"/issues/"+strconv.Itoa(id)+"/watchers.json?key="+key, "application/json", watcherJson)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteIssue(url, key string, id, timeout int, customClient *http.Client) error {
	var client *utils.Api
	if customClient == nil {
		client = utils.NewHttpClient(timeout)
	} else {
		client = utils.NewHttpClient(timeout, utils.OptionHTTPClient(customClient))
	}
	err := client.Delete(url + "/issues/" + strconv.Itoa(id) + ".json?key=" + key)
	if err != nil {
		return err
	}
	return nil
}
