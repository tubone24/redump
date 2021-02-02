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
	conf, err := config.GetConfig()
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
				uploadFile, err := redmine.UploadAttachmentFiles(conf.NewServerConfig.Url, conf.NewServerConfig.Key, fileParams)
				if err != nil {
					panic(err)
				}
				uploadFiles = append(uploadFiles, uploadFile[0])
			}
		}
		convertedIssueParam := redmine.CreateIssueParam(*convertedIssue, uploadFiles)
		issueId, err := redmine.CreateIssue(conf.NewServerConfig.Url, conf.NewServerConfig.Key, convertedIssueParam)
		if err != nil {
			return err
		}
		notes := redmine.CreateJournalStrings(*convertedIssue)
		err = redmine.UpdateIssueJournals(conf.NewServerConfig.Url, conf.NewServerConfig.Key, issueId, notes)
		if err != nil {
			return err
		}
		err = redmine.UpdateWatchers(conf.NewServerConfig.Url, conf.NewServerConfig.Key, issueId, *convertedIssue)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * time.Duration(conf.ServerConfig.Sleep))
	}
	return nil
}
