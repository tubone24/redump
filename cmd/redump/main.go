package main

import (
	"fmt"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/utils"
	"strconv"
	"sync"
)

func main(){
	issues, err := redmine.GetIssues("http://localhost:8889", "")
	var wg sync.WaitGroup
	if err != nil {
		panic(err)
	}
	for _, v := range issues {
		wg.Add(1)
		go func(id int){
			data, err := redmine.GetIssue("http://localhost:8889", "", id)
			if err != nil {
				panic(err)
			}
			bolB, _ := json.Marshal(data)
			err = utils.WriteFile(strconv.Itoa(id) + ".json", bolB)
			if err != nil {
			}
			fmt.Println("Write: " + strconv.Itoa(id) + ".json")
			wg.Done()
		}(v.Id)
	}
	wg.Wait()
}

func run(wg sync.WaitGroup, issue *redmine.Issue) {
	data, err := redmine.GetIssue("http://localhost:8889", "", issue.Id)
	if err != nil {
		panic(err)
	}
	bolB, _ := json.Marshal(data)
	err = utils.WriteFile(strconv.Itoa(issue.Id) + ".json", bolB)
	if err != nil {
	}
	wg.Done()
}
