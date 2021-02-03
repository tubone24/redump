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
	cfg, err := config.GetConfig("")
	if err != nil {
		panic(err)
	}
	issues, err := redmine.GetIssues(cfg.ServerConfig.Url, cfg.ServerConfig.Key, projectId, cfg.ServerConfig.Timeout, nil)
	if err != nil {
		panic(err)
	}
	txtCh := make(chan string, 10)
	var wg sync.WaitGroup
	defer close(txtCh)
	if concurrency {
		for _, v := range issues {
			wg.Add(1)
			go func(issue redmine.Issue, conf config.Config) {
				detailIssue, err := redmine.GetIssue(conf.ServerConfig.Url, conf.ServerConfig.Key, issue.Id, cfg.ServerConfig.Timeout, nil)
				if err != nil {
					panic(err)
				}
				issueJson, _ := json.Marshal(detailIssue)
				err = utils.WriteFile("data/issues/"+strconv.Itoa(issue.Id)+".json", issueJson)
				if detailIssue.Attachments != nil {
					downloadBody, err := redmine.DownloadAttachmentFiles(conf.ServerConfig.Key, cfg.ServerConfig.Timeout, detailIssue.Attachments, nil)
					if err != nil {
						panic(err)
					}
					for index, file := range downloadBody {
						err = utils.WriteFile("data/issues/attachments/"+strconv.Itoa(issue.Id)+"_"+strconv.Itoa(index)+"_"+detailIssue.Attachments[index].FileName, file)
						if err != nil {
							panic(err)
						}
					}
				}
				wg.Done()
				fmt.Println("Success: " + strconv.Itoa(issue.Id) + ".json")
			}(*v, *cfg)
		}
		wg.Wait()
	} else {
		for _, v := range issues {
			go runDump(txtCh, v.Id)
			time.Sleep(time.Millisecond * time.Duration(cfg.ServerConfig.Sleep))
			fmt.Println(<-txtCh)
		}
	}
}

func DumpOneIssue(issueId int) {
	txtCh := make(chan string, 10)
	defer close(txtCh)
	go runDump(txtCh, issueId)
	fmt.Println(<-txtCh)
}

func runDump(txtCh chan<- string, issueId int) {
	cfg, err := config.GetConfig("")
	if err != nil {
		panic(err)
	}
	detailIssue, err := redmine.GetIssue(cfg.ServerConfig.Url, cfg.ServerConfig.Key, issueId, cfg.ServerConfig.Timeout, nil)
	if err != nil {
		panic(err)
	}
	issueJson, _ := json.Marshal(detailIssue)
	err = utils.WriteFile("data/issues/"+strconv.Itoa(issueId)+".json", issueJson)
	if detailIssue.Attachments != nil {
		downloadBody, err := redmine.DownloadAttachmentFiles(cfg.ServerConfig.Key, cfg.ServerConfig.Timeout, detailIssue.Attachments, nil)
		if err != nil {
			panic(err)
		}
		for index, file := range downloadBody {
			err = utils.WriteFile("data/issues/attachments/"+strconv.Itoa(issueId)+"_"+strconv.Itoa(index)+"_"+detailIssue.Attachments[index].FileName, file)
			if err != nil {
				panic(err)
			}
		}
	}
	if err != nil {
		txtCh <- "Failed: " + strconv.Itoa(issueId) + ".json"
	}
	txtCh <- "Success: " + strconv.Itoa(issueId) + ".json"
}
