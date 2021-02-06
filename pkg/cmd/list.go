package cmd

import (
	"github.com/tubone24/redump/pkg/config"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/tubone24/redump/pkg/utils"
	"net/http"
)

func ListAll(projectId int) error {
	cfg, err := config.GetConfig("")
	if err != nil {
		return err
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
	_, err = redmine.ListProjectId(issues, "")
	if err != nil {
		return err
	}
	_, err = redmine.ListTrackerId(issues, "")
	if err != nil {
		return err
	}
	_, err = redmine.ListStatusId(issues, "")
	if err != nil {
		return err
	}
	_, err = redmine.ListPriorityId(issues, "")
	if err != nil {
		return err
	}
	_, err = redmine.ListUserIdAssignedTo(issues, "")
	if err != nil {
		return err
	}
	_, err = redmine.ListCustomFieldsId(issues, "")
	if err != nil {
		return err
	}
	return nil
}
