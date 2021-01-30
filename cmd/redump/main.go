package main

import (
	"fmt"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/utils"
	"github.com/tubone24/redump/pkg/config"
	"strconv"
)

func main(){
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	issueParam := redmine.IssueParam{ProjectId: 1, TrackerId: 1, Subject: "test2", PriorityId: 1, StatusId: 1}
	id, err := redmine.CreateIssue(config.ServerConfig.Url, config.ServerConfig.Key, issueParam)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	err = redmine.UpdateIssueJournals(config.ServerConfig.Url, config.ServerConfig.Key, id, [] string{"Golang", "Java"})
	if err != nil {
		panic(err)
	}
	issues, err := redmine.GetIssues(config.ServerConfig.Url, config.ServerConfig.Key)
	txtCh := make(chan string, 10)
	defer close(txtCh)
	if err != nil {
		panic(err)
	}
	for _, v := range issues {
		go run(txtCh, v)
		fmt.Println(<-txtCh)
	}
}

func run(txtCh chan<- string, issue *redmine.Issue) {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	data, err := redmine.GetIssue(config.ServerConfig.Url, config.ServerConfig.Key, issue.Id)
	if err != nil {
		panic(err)
	}
	bolB, _ := json.Marshal(data)
	err = utils.WriteFile("data/" + strconv.Itoa(issue.Id) + ".json", bolB)
	if err != nil {
		txtCh <- "Failed: " + strconv.Itoa(issue.Id) + ".json"
	}
	txtCh <- "Success: " + strconv.Itoa(issue.Id) + ".json"
}
