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
	_, err = redmine.ListProjectId(issues, "data/project_id.json")
	if err != nil {
		return err
	}
	_, err = redmine.ListTrackerId(issues, "data/tracker_id.json")
	if err != nil {
		return err
	}
	_, err = redmine.ListStatusId(issues, "data/status_id.json")
	if err != nil {
		return err
	}
	_, err = redmine.ListPriorityId(issues, "data/priority_id.json")
	if err != nil {
		return err
	}
	_, err = redmine.ListUserIdAssignedTo(issues, "data/userId.json")
	if err != nil {
		return err
	}
	_, err = redmine.ListCustomFieldsId(issues, "data/custom_fields_id.json")
	if err != nil {
		return err
	}
	return nil
}
