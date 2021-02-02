package utils_test

import (
	"github.com/tubone24/redump/pkg/utils"
	"testing"
)

func TestGetContentType(t *testing.T) {
	patterns := []struct {
		key      string
		expected string
	}{
		{"test.txt", "text/plain"},
		{"test.tsv", "text/tab-separated-values"},
		{"test.csv", "text/csv"},
		{"test.json", "application/json"},
		{"test.pdf", "application/pdf"},
		{"test.xls", "application/vnd.ms-excel"},
		{"test.xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
		{"test.ppt", "application/vnd.ms-powerpoint"},
		{"test.pptx", "application/vnd.openxmlformats-officedocument.presentationml.presentation"},
		{"test.doc", "application/msword"},
		{"test.docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document"},
		{"test.zip", "application/zip"},
		{"test.lzh", "application/x-lzh"},
		{"test.tar.gz", "application/x-tar"},
		{"test.tgz", "application/x-tar"},
		{"test.tar", "application/x-tar"},
		{"test.bz", "application/x-bzip"},
		{"test.bz2", "application/x-bzip2"},
		{"test.gz", "application/gzip"},
		{"test.rar", "application/vnd.rar"},
		{"test.7z", "application/x-7z-compressed"},
		{"test.xml", "application/xml"},
		{"test.bin", "application/octet-stream"},
		{"test.mse", "application/x-www-form-urlencoded"},
	}

	for idx, pattern := range patterns {
		actual := utils.GetContentType(pattern.key)
		if pattern.expected != actual {
			t.Errorf("testcase%d: expected: %s, actual %s", idx, pattern.expected, actual)
		}
	}
}
