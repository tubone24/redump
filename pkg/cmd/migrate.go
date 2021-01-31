package cmd

import (
	"fmt"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/utils"
	"github.com/tubone24/redump/pkg/config"
	"strconv"
)

func Migrate(projectId int) {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	issues, err := redmine.GetIssues(config.ServerConfig.Url, config.ServerConfig.Key, projectId)
	txtCh := make(chan string, 10)
	defer close(txtCh)
	if err != nil {
		panic(err)
	}
	for _, v := range issues {
		go runMigrate(txtCh, v)
		fmt.Println(<-txtCh)
	}
}

func runMigrate(txtCh chan<- string, issue *redmine.Issue) {
	var uploadFiles []redmine.FileParam
	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	detailIssue, err := redmine.GetIssue(conf.ServerConfig.Url, conf.ServerConfig.Key, issue.Id)
	if err != nil {
		panic(err)
	}
	issueJson, _ := json.Marshal(detailIssue)
	err = utils.WriteFile("data/"+strconv.Itoa(issue.Id)+".json", issueJson)
	if detailIssue.Attachments != nil {
		downloadBody, err := redmine.DownloadAttachmentFiles(conf.ServerConfig.Key, detailIssue.Attachments)
		if err != nil {
			panic(err)
		}
		for index, file := range downloadBody {
			err = utils.WriteFile("data/"+strconv.Itoa(issue.Id)+"_"+strconv.Itoa(index)+"_"+detailIssue.Attachments[index].FileName, file)
			if err != nil {
				panic(err)
			}
			fileParam := redmine.FileParam{FileName: detailIssue.Attachments[index].FileName, ContentType: utils.GetContentType(detailIssue.Attachments[index].FileName), Contents: file}
			fileParams := []redmine.FileParam{fileParam}
			uploadFile, err := redmine.UploadAttachmentFiles(conf.ServerConfig.Url, conf.ServerConfig.Key, fileParams)
			fmt.Println(uploadFile[0].Token)
			if err != nil {
				panic(err)
			}
			uploadFiles = append(uploadFiles, uploadFile[0])
		}
	}
	newIssue, err := redmine.ConvertNewEnv(detailIssue)
	if err != nil {
		panic(err)
	}
	var newIssueParam redmine.IssueParam
	if newIssue.Attachments != nil {
		var uploads []redmine.Uploads
		for _, v := range uploadFiles {
			uploads = append(uploads, redmine.Uploads{FileName: v.FileName, ContentType: v.ContentType, Token: v.Token})
		}
		newIssueParam = redmine.IssueParam{
			ProjectId: newIssue.Project.Id,
			TrackerId: newIssue.Tracker.Id,
			StatusId: newIssue.Status.Id,
			PriorityId: newIssue.Priority.Id,
			AssignedToId: newIssue.AssignedTo.Id,
			Subject:newIssue.Subject,
			Description: newIssue.Description,
			CustomFields: newIssue.CustomFields,
			Uploads: uploads}
	} else {
		newIssueParam = redmine.IssueParam{
			ProjectId: newIssue.Project.Id,
			TrackerId: newIssue.Tracker.Id,
			StatusId: newIssue.Status.Id,
			PriorityId: newIssue.Priority.Id,
			AssignedToId: newIssue.AssignedTo.Id,
			Subject:newIssue.Subject,
			Description: newIssue.Description,
			CustomFields: newIssue.CustomFields}
	}
	issueId, err := redmine.CreateIssue(conf.ServerConfig.Url, conf.ServerConfig.Key, newIssueParam)
	if err != nil {
		panic(err)
	}
	notes := redmine.CreateJournalStrings(*newIssue)
	err = redmine.UpdateIssueJournals(conf.ServerConfig.Url, conf.ServerConfig.Key, issueId, notes)
	if err != nil {
		txtCh <- "Failed: " + strconv.Itoa(issue.Id) + ".json"
	}
	txtCh <- "Success: " + strconv.Itoa(issue.Id) + ".json"
}
