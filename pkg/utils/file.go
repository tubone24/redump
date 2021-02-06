package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func WriteFile(filename string, output []byte) error {
	file, err := os.Create(filename)
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

func CheckDir(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func MakeDir(dir string) error {
	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}
	return nil
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
	case ".js", ".mjs":
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
	case ".lzh", ".lha":
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
	case ".bin", ".class", ".dll", ".dmg", ".dms", ".dpc", ".dpt", ".dpv", ".exe", ".mvr", ".so", ".wpi", ".cab":
		return "application/octet-stream"
	case ".aac":
		return "audio/aac"
	case ".abw":
		return "application/x-abiword"
	case ".arc":
		return "application/x-freearc"
	case ".avi":
		return "video/x-msvideo"
	case ".azw":
		return "application/vnd.amazon.ebook"
	case ".csh":
		return "application/x-csh"
	case ".eot":
		return "application/vnd.ms-fontobject"
	case ".epub":
		return "application/epub+zip"
	case ".ico":
		return "image/vnd.microsoft.icon"
	case ".ics":
		return "text/calendar"
	case ".jar":
		return "application/java-archive"
	case ".jsonld":
		return "application/ld+json"
	case ".mid", ".midi":
		return "audio/midi"
	case ".mpkg":
		return "application/vnd.apple.installer+xml"
	case ".odp":
		return "application/vnd.oasis.opendocument.presentation"
	case ".ods":
		return "application/vnd.oasis.opendocument.spreadsheet"
	case ".odt":
		return "application/vnd.oasis.opendocument.text"
	case ".oga":
		return "audio/ogg"
	case ".ogv":
		return "video/ogg"
	case ".ogx":
		return "application/ogg"
	case ".opus":
		return "audio/opus"
	case ".otf":
		return "font/otf"
	case ".php":
		return "application/x-httpd-php"
	case ".rtf":
		return "application/rtf"
	case ".sh":
		return "application/x-sh"
	case ".swf":
		return "application/x-shockwave-flash"
	case ".tif", ".tiff":
		return "image/tiff"
	case ".ts":
		return "video/mp2t"
	case ".ttf":
		return "font/ttf"
	case ".vsd":
		return "application/vnd.visio"
	case ".wav":
		return "audio/wav"
	case ".weba":
		return "audio/webm"
	case ".webm":
		return "video/webm"
	case ".webp":
		return "image/webp"
	case ".woff":
		return "font/woff"
	case ".woff2":
		return "font/woff2"
	case ".xhtml":
		return "application/xhtml+xml"
	case ".xul":
		return "application/vnd.mozilla.xul+xml"
	case ".3gp":
		return "video/3gpp"
	case ".3g2":
		return "video/3gpp2"
	case ".jtd":
		return "application/x-js-taro"
	case ".mov", ".qt":
		return "video/quicktime"
	case ".vcf":
		return "text/x-vcard"
	case ".vcs":
		return "text/x-vcalendar"
	case ".kml":
		return "application/vnd.google-earth.kml+xml"
	case ".kmz":
		return "application/vnd.google-earth.kmz"
	case ".latex":
		return "application/x-latex"
	case ".tex":
		return "application/x-tex"
	case ".dmt":
		return "application/x-decomail-template"
	case ".atom":
		return "application/atom+xml"
	case ".wmv":
		return "video/x-ms-wmv"
	case ".flv":
		return "video/x-flv"
	default:
		return "application/x-www-form-urlencoded"
	}
}

func SanitizeInvalidFileName(str string) (result string) {
	fmt.Println(str)
	result = strings.Replace(strings.Replace(str, " ", "_", -1), "　", "_", -1)
	fmt.Println(result)
	return
}
