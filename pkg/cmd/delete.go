package cmd

import (
	"bufio"
	"fmt"
	"github.com/tubone24/redump/pkg/config"
	"github.com/tubone24/redump/pkg/redmine"
	"os"
	"sync"
)

//socket: too many open files
func DeleteServerAllIssuesConcurrency(old bool) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}
	var serverUrl string
	var serverKey string
	if old {
		serverUrl = cfg.ServerConfig.Url
		serverKey = cfg.ServerConfig.Key
	} else {
		serverUrl = cfg.NewServerConfig.Url
		serverKey = cfg.NewServerConfig.Key
	}
	fmt.Print("Are you sure delete all issues on " + serverUrl + " ?(y/n)=>")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Text() == "y" {
		var wg sync.WaitGroup
		issues, err := redmine.GetIssues(serverUrl, serverKey, 0, cfg.ServerConfig.Timeout)
		if err != nil {
			return err
		}
		for _, v := range issues {
			go func(issue *redmine.Issue) {
				wg.Add(1)
				err := redmine.DeleteIssue(serverUrl, serverKey, issue.Id, cfg.ServerConfig.Timeout)
				fmt.Println(issue.Id)
				if err != nil {
					panic(err)
				}
				wg.Done()
			}(v)
		}
		wg.Wait()
		return nil
	}
	return nil
}

func DeleteServerAllIssues(old bool) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}
	var serverUrl string
	var serverKey string
	if old {
		serverUrl = cfg.ServerConfig.Url
		serverKey = cfg.ServerConfig.Key
	} else {
		serverUrl = cfg.NewServerConfig.Url
		serverKey = cfg.NewServerConfig.Key
	}
	fmt.Print("Are you sure delete all issues on " + serverUrl + " ?(y/n)=>")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Text() == "y" {
		issues, err := redmine.GetIssues(serverUrl, serverKey, 0, cfg.ServerConfig.Timeout)
		if err != nil {
			return err
		}
		for _, v := range issues {
			err := redmine.DeleteIssue(serverUrl, serverKey, v.Id, cfg.ServerConfig.Timeout)
			fmt.Println(v.Id)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}
