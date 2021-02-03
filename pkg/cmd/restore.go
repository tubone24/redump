package cmd

import (
	"github.com/tubone24/redump/pkg/config"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/tubone24/redump/pkg/utils"
	"path/filepath"
	"strconv"
	"time"
)

func RestoreDataFromLocal(projectId, issueId int) error {
	cfg, err := config.GetConfig("")
	if err != nil {
		return err
	}
	var jsonFiles []string
	if issueId != 0 {
		jsonFiles, _ = filepath.Glob("data/issues/" + strconv.Itoa(issueId) + ".json")
	} else {
		jsonFiles, _ = filepath.Glob("data/issues/*.json")
	}
	for _, file := range jsonFiles {
		var uploadFiles []redmine.FileParam
		issueJsonBytes, err := utils.ReadFile(file)
		if err != nil {
			return err
		}
		issue, err := redmine.CreateIssueFromByteSlice(issueJsonBytes)
		if err != nil {
			return err
		}
		if issue.Project.Id != projectId {
			continue
		}
		convertedIssue, err := redmine.ConvertNewEnv(*issue)
		if err != nil {
			return err
		}
		if convertedIssue.Attachments != nil {
			for index, attachment := range convertedIssue.Attachments {
				filename := "data/issues/attachments/" + strconv.Itoa(convertedIssue.Id) + "_" + strconv.Itoa(index) + "_" + attachment.FileName
				contentBytes, err := utils.ReadFile(filename)
				if err != nil {
					return err
				}
				fileParam := redmine.FileParam{FileName: convertedIssue.Attachments[index].FileName, ContentType: utils.GetContentType(convertedIssue.Attachments[index].FileName), Contents: contentBytes}
				fileParams := []redmine.FileParam{fileParam}
				uploadFile, err := redmine.UploadAttachmentFiles(cfg.NewServerConfig.Url, cfg.NewServerConfig.Key, cfg.ServerConfig.Timeout, fileParams, nil)
				if err != nil {
					panic(err)
				}
				uploadFiles = append(uploadFiles, uploadFile[0])
			}
		}
		convertedIssueParam := redmine.CreateIssueParam(*convertedIssue, uploadFiles)
		issueId, err := redmine.CreateIssue(cfg.NewServerConfig.Url, cfg.NewServerConfig.Key, cfg.ServerConfig.Timeout, convertedIssueParam, nil)
		if err != nil {
			return err
		}
		notes := redmine.CreateJournalStrings(*convertedIssue)
		err = redmine.UpdateIssueJournals(cfg.NewServerConfig.Url, cfg.NewServerConfig.Key, issueId, cfg.ServerConfig.Timeout, notes, nil)
		if err != nil {
			return err
		}
		err = redmine.UpdateWatchers(cfg.NewServerConfig.Url, cfg.NewServerConfig.Key, issueId, cfg.ServerConfig.Timeout, *convertedIssue, nil)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * time.Duration(cfg.ServerConfig.Sleep))
	}
	return nil
}
