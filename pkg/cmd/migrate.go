package cmd

import (
	"fmt"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/utils"
	"github.com/tubone24/redump/pkg/config"
	"strconv"
	"time"
)

func Migrate(projectId int) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}
	issues, err := redmine.GetIssues(cfg.ServerConfig.Url, cfg.ServerConfig.Key, projectId)
	txtCh := make(chan string, 10)
	defer close(txtCh)
	if err != nil {
		return err
	}
	for _, v := range issues {
		go runMigrateIssue(txtCh, v.Id)
		time.Sleep(time.Millisecond * time.Duration(cfg.ServerConfig.Sleep))
		fmt.Println(<-txtCh)
	}
	return nil
}

func MigrateOneIssue(issueId int) {
	txtCh := make(chan string)
	defer close(txtCh)
	go runMigrateIssue(txtCh, issueId)
	fmt.Println(<-txtCh)
}

func runMigrateIssue(txtCh chan<- string, issueId int) {
	var uploadFiles []redmine.FileParam
	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	detailIssue, err := redmine.GetIssue(conf.ServerConfig.Url, conf.ServerConfig.Key, issueId)
	if err != nil {
		panic(err)
	}
	issueJson, _ := json.Marshal(detailIssue)
	err = utils.WriteFile("data/issues/"+strconv.Itoa(issueId)+".json", issueJson)
	if detailIssue.Attachments != nil {
		downloadBody, err := redmine.DownloadAttachmentFiles(conf.ServerConfig.Key, detailIssue.Attachments)
		if err != nil {
			panic(err)
		}
		for index, file := range downloadBody {
			err = utils.WriteFile("data/issues/attachments/"+strconv.Itoa(issueId)+"_"+strconv.Itoa(index)+"_"+detailIssue.Attachments[index].FileName, file)
			if err != nil {
				panic(err)
			}
			fileParam := redmine.FileParam{FileName: detailIssue.Attachments[index].FileName, ContentType: utils.GetContentType(detailIssue.Attachments[index].FileName), Contents: file}
			fileParams := []redmine.FileParam{fileParam}
			uploadFile, err := redmine.UploadAttachmentFiles(conf.NewServerConfig.Url, conf.NewServerConfig.Key, fileParams)
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
	newIssueParam := redmine.CreateIssueParam(*newIssue, uploadFiles)
	newIssueId, err := redmine.CreateIssue(conf.NewServerConfig.Url, conf.NewServerConfig.Key, newIssueParam)
	if err != nil {
		panic(err)
	}
	notes := redmine.CreateJournalStrings(*newIssue)
	err = redmine.UpdateIssueJournals(conf.NewServerConfig.Url, conf.NewServerConfig.Key, newIssueId, notes)
	if err != nil {
		panic(err)
	}
	err = redmine.UpdateWatchers(conf.NewServerConfig.Url, conf.NewServerConfig.Key, newIssueId, *newIssue)
	if err != nil {
		txtCh <- "Failed: " + strconv.Itoa(issueId) + ".json"
	}
	txtCh <- "Success: " + strconv.Itoa(issueId) + ".json"
}
