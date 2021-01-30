package main

import (
	"fmt"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/utils"
	"strconv"
)

func main(){
	issueParam := redmine.IssueParam{ProjectId: 1, TrackerId: 1, Subject: "test", PriorityId: 1, StatusId: 1, AssignedToId:1}
	err := redmine.CreateIssue("http://localhost:8889", "573f9544d1cc7512d4eed751b1d79d23210e964b", issueParam)
	if err != nil {
		panic(err)
	}
	issues, err := redmine.GetIssues("http://localhost:8889", "573f9544d1cc7512d4eed751b1d79d23210e964b")
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
	data, err := redmine.GetIssue("http://localhost:8889", "573f9544d1cc7512d4eed751b1d79d23210e964b", issue.Id)
	if err != nil {
		panic(err)
	}
	bolB, _ := json.Marshal(data)
	err = utils.WriteFile(strconv.Itoa(issue.Id) + ".json", bolB)
	if err != nil {
		txtCh <- "Failed: " + strconv.Itoa(issue.Id) + ".json"
	}
	txtCh <- "Success: " + strconv.Itoa(issue.Id) + ".json"
}
