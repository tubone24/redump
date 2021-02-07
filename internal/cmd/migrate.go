package cmd

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/config"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/tubone24/redump/pkg/utils"
	"net/http"
	"strconv"
	"time"
)

func Migrate(projectId int, silent bool) error {
	cfg, err := config.GetConfig("")
	if err != nil {
		return err
	}
	if !utils.CheckDir("data/issues/attachments") {
		err := utils.MakeDir("data/issues/attachments")
		if err != nil {
			panic(err)
		}
	}
	var customClient *http.Client
	if cfg.ServerConfig.ProxyUrl != "" {
		customClient, err = utils.NewProxyClient(cfg.ServerConfig.ProxyUrl)
		if err != nil {
			panic(err)
		}
	} else {
		customClient = nil
	}
	issues, err := redmine.GetIssues(cfg.ServerConfig.Url, cfg.ServerConfig.Key, projectId, cfg.ServerConfig.Timeout, customClient)
	txtCh := make(chan string, 10)
	defer close(txtCh)
	if err != nil {
		return err
	}
	for _, v := range issues {
		go runMigrateIssue(txtCh, v.Id, silent)
		time.Sleep(time.Millisecond * time.Duration(cfg.ServerConfig.Sleep))
		fmt.Println(<-txtCh)
	}
	return nil
}

func MigrateOneIssue(issueId int, silent bool) {
	if !utils.CheckDir("data/issues/attachments") {
		err := utils.MakeDir("data/issues/attachments")
		if err != nil {
			panic(err)
		}
	}
	txtCh := make(chan string)
	defer close(txtCh)
	go runMigrateIssue(txtCh, issueId, silent)
	fmt.Println(<-txtCh)
}

func runMigrateIssue(txtCh chan<- string, issueId int, silent bool) {
	var uploadFiles []redmine.FileParam
	cfg, err := config.GetConfig("")
	if err != nil {
		panic(err)
	}
	var oldCustomClient *http.Client
	if cfg.ServerConfig.ProxyUrl != "" {
		oldCustomClient, err = utils.NewProxyClient(cfg.ServerConfig.ProxyUrl)
		if err != nil {
			panic(err)
		}
	} else {
		oldCustomClient = nil
	}
	var newCustomClient *http.Client
	if cfg.NewServerConfig.ProxyUrl != "" {
		oldCustomClient, err = utils.NewProxyClient(cfg.NewServerConfig.ProxyUrl)
		if err != nil {
			panic(err)
		}
	} else {
		oldCustomClient = nil
	}
	detailIssue, err := redmine.GetIssue(cfg.ServerConfig.Url, cfg.ServerConfig.Key, issueId, cfg.ServerConfig.Timeout, oldCustomClient)
	if err != nil {
		panic(err)
	}
	issueJson, _ := json.Marshal(detailIssue)
	err = utils.WriteFile("data/issues/"+strconv.Itoa(issueId)+".json", issueJson)
	if detailIssue.Attachments != nil {
		downloadBody, err := redmine.DownloadAttachmentFiles(cfg.ServerConfig.Key, cfg.ServerConfig.Timeout, detailIssue.Attachments, oldCustomClient)
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
			uploadFile, err := redmine.UploadAttachmentFiles(cfg.NewServerConfig.Url, cfg.NewServerConfig.Key, cfg.NewServerConfig.Timeout, fileParams, newCustomClient)
			fmt.Println(uploadFile[0].Token)
			if err != nil {
				panic(err)
			}
			uploadFiles = append(uploadFiles, uploadFile[0])
		}
	}
	newIssue, err := redmine.ConvertNewEnv(detailIssue, *cfg, silent)
	if err != nil {
		panic(err)
	}
	newIssueParam := redmine.CreateIssueParam(*newIssue, uploadFiles)
	newIssueId, err := redmine.CreateIssue(cfg.NewServerConfig.Url, cfg.NewServerConfig.Key, cfg.NewServerConfig.Timeout, newIssueParam, newCustomClient)
	if err != nil {
		panic(err)
	}
	notes := redmine.CreateJournalStrings(*newIssue)
	err = redmine.UpdateIssueJournals(cfg.NewServerConfig.Url, cfg.NewServerConfig.Key, newIssueId, cfg.NewServerConfig.Timeout, notes, newCustomClient)
	if err != nil {
		panic(err)
	}
	err = redmine.UpdateWatchers(cfg.NewServerConfig.Url, cfg.NewServerConfig.Key, newIssueId, cfg.NewServerConfig.Timeout, *newIssue, newCustomClient)
	if err != nil {
		txtCh <- "Failed: " + strconv.Itoa(issueId) + ".json"
	}
	txtCh <- "Success: " + strconv.Itoa(issueId) + ".json"
}
