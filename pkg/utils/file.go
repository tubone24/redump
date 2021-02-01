package utils

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func WriteFile(filename string, output []byte) error {
	file, err  := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(output)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(file string) ([]byte, error) {
	const bufferSize = 256
	var content []byte
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	buffer := make([]byte, bufferSize)
	for {
		n, err := fp.Read(buffer)
		if 0 < n {
			content = append(content, buffer...)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return content, nil
}

func GetContentType(key string) string {
	if strings.HasSuffix(key, ".tar.gz") {
		return "application/x-tar"
	}

	ext := filepath.Ext(key)

	switch ext {
	case ".txt":
		return "text/plain"
	case ".htm", ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "text/javascript"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".bmp":
		return "image/bmp"
	case ".svg":
		return "image/svg+xml"
	case ".mp3":
		return "audio/mpeg"
	case ".mp4":
		return "video/mp4"
	case ".mpg", ".mpeg":
		return "video/mpeg"
	case ".tsv":
		return "text/tab-separated-values"
	case ".csv":
		return "text/csv"
	case ".json":
		return "application/json"
	case ".pdf":
		return "application/pdf"
	case ".xls":
		return "application/vnd.ms-excel"
	case ".xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case ".ppt":
		return "application/vnd.ms-powerpoint"
	case ".pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case ".doc":
		return "application/msword"
	case ".docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case ".zip":
		return "application/zip"
	case ".lzh":
		return "application/x-lzh"
	case ".tar.gz":
		return "application/x-tar"
	case ".tgz":
		return "application/x-tar"
	case ".tar":
		return "application/x-tar"
	case ".bz":
		return "application/x-bzip"
	case ".bz2":
		return "application/x-bzip2"
	case ".gz":
		return "application/gzip"
	case ".rar":
		return "application/vnd.rar"
	case ".7z":
		return "application/x-7z-compressed"
	case ".xml":
		return "application/xml"
	case ".bin":
		return "application/octet-stream"
	default:
		return "application/x-www-form-urlencoded"
	}
}
