package cmd

import (
	"fmt"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/utils"
	"github.com/tubone24/redump/pkg/config"
	"strconv"
	"sync"
	"time"
)

func Dump(projectId int, concurrency bool) {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	issues, err := redmine.GetIssues(config.ServerConfig.Url, config.ServerConfig.Key, projectId)
	txtCh := make(chan string, 10)
	var wg sync.WaitGroup
	defer close(txtCh)
	if err != nil {
		panic(err)
	}
	if concurrency {
		for _, v := range issues {
			wg.Add(1)
			go runDumpConcurrency(v, wg)
		}
		wg.Wait()
	} else {
		for _, v := range issues {
			go runDump(txtCh, v)
			time.Sleep(time.Millisecond * time.Duration(config.ServerConfig.Sleep))
			fmt.Println(<-txtCh)
		}
	}
}

func runDump(txtCh chan<- string, issue *redmine.Issue) {
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
		}
	}
	if err != nil {
		txtCh <- "Failed: " + strconv.Itoa(issue.Id) + ".json"
	}
	txtCh <- "Success: " + strconv.Itoa(issue.Id) + ".json"
}

func runDumpConcurrency(issue *redmine.Issue, wg sync.WaitGroup) {
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
		}
	}
	wg.Done()
}
