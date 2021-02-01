package cmd

import (
	"path/filepath"
	"github.com/tubone24/redump/pkg/redmine"
	"github.com/tubone24/redump/pkg/utils"
	"github.com/tubone24/redump/pkg/config"
	"strconv"
)

func RestoreDataFromLocal() error {
	var uploadFiles []redmine.FileParam
	conf, err := config.GetConfig()
	if err != nil {
		return err
	}
	jsonFiles, _ := filepath.Glob("data/*.json")
	for _, file := range jsonFiles {
		issueJsonBytes, err := utils.ReadFile(file)
		if err != nil {
			return err
		}
		issue, err := redmine.CreateIssueFromByteSlice(issueJsonBytes)
		if err != nil {
			return err
		}
		convertedIssue, err := redmine.ConvertNewEnv(*issue)
		if err != nil {
			return err
		}
		if convertedIssue.Attachments != nil {
			for index, attachment := range convertedIssue.Attachments {
				filename := strconv.Itoa(convertedIssue.Id) + "_" + strconv.Itoa(index) + "_" + attachment.FileName
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
	}
	return nil
}