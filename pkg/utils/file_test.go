package utils_test

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/tubone24/redump/pkg/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGetContentType(t *testing.T) {
	patterns := []struct {
		key      string
		expected string
	}{
		{"test.txt", "text/plain"},
		{"test.htm", "text/html"},
		{"test.html", "text/html"},
		{"test.css", "text/css"},
		{"test.js", "text/javascript"},
		{"test.mjs", "text/javascript"},
		{"test.png", "image/png"},
		{"test.jpg", "image/jpeg"},
		{"test.jpeg", "image/jpeg"},
		{"test.gif", "image/gif"},
		{"test.bmp", "image/bmp"},
		{"test.svg", "image/svg+xml"},
		{"test.mp3", "audio/mpeg"},
		{"test.mp4", "video/mp4"},
		{"test.mpg", "video/mpeg"},
		{"test.mpeg", "video/mpeg"},
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
		{"test.class", "application/octet-stream"},
		{"test.dll", "application/octet-stream"},
		{"test.dmg", "application/octet-stream"},
		{"test.dms", "application/octet-stream"},
		{"test.dpc", "application/octet-stream"},
		{"test.dpt", "application/octet-stream"},
		{"test.dpv", "application/octet-stream"},
		{"test.exe", "application/octet-stream"},
		{"test.mvr", "application/octet-stream"},
		{"test.so", "application/octet-stream"},
		{"test.wpi", "application/octet-stream"},
		{"test.cab", "application/octet-stream"},
		{"test.aac", "audio/aac"},
		{"test.abw", "application/x-abiword"},
		{"test.arc", "application/x-freearc"},
		{"test.avi", "video/x-msvideo"},
		{"test.azw", "application/vnd.amazon.ebook"},
		{"test.csh", "application/x-csh"},
		{"test.eot", "application/vnd.ms-fontobject"},
		{"test.epub", "application/epub+zip"},
		{"test.ico", "image/vnd.microsoft.icon"},
		{"test.ics", "text/calendar"},
		{"test.jsonld", "application/ld+json"},
		{"test.mid", "audio/midi"},
		{"test.midi", "audio/midi"},
		{"test.mpkg", "application/vnd.apple.installer+xml"},
		{"test.odp", "application/vnd.oasis.opendocument.presentation"},
		{"test.ods", "application/vnd.oasis.opendocument.spreadsheet"},
		{"test.odt", "application/vnd.oasis.opendocument.text"},
		{"test.oga", "audio/ogg"},
		{"test.ogv", "video/ogg"},
		{"test.ogx", "application/ogg"},
		{"test.opus", "audio/opus"},
		{"test.otf", "font/otf"},
		{"test.php", "application/x-httpd-php"},
		{"test.rtf", "application/rtf"},
		{"test.tif", "image/tiff"},
		{"test.tiff", "image/tiff"},
		{"test.ts", "video/mp2t"},
		{"test.ttf", "font/ttf"},
		{"test.vsd", "application/vnd.visio"},
		{"test.wav", "audio/wav"},
		{"test.weba", "audio/webm"},
		{"test.webm", "video/webm"},
		{"test.webp", "image/webp"},
		{"test.woff", "font/woff"},
		{"test.woff2", "font/woff2"},
		{"test.xhtml", "application/xhtml+xml"},
		{"test.xul", "application/vnd.mozilla.xul+xml"},
		{"test.3gp", "video/3gpp"},
		{"test.3g2", "video/3gpp2"},
		{"test.jtd", "application/x-js-taro"},
		{"test.mov", "video/quicktime"},
		{"test.qt", "video/quicktime"},
		{"test.vcf", "text/x-vcard"},
		{"test.vcs", "text/x-vcalendar"},
		{"test.kml", "application/vnd.google-earth.kml+xml"},
		{"test.kmz", "application/vnd.google-earth.kmz"},
		{"test.latex", "application/x-latex"},
		{"test.tex", "application/x-tex"},
		{"test.dmt", "application/x-decomail-template"},
		{"test.atom", "application/atom+xml"},
		{"test.wmv", "video/x-ms-wmv"},
		{"test.flv", "video/x-flv"},
		{"test.swf", "application/x-shockwave-flash"},
		{"test.tar.gz", "application/x-tar"},
		{"test.jar", "application/java-archive"},
		{"test.sh", "application/x-sh"},
		{"test.mse", "application/x-www-form-urlencoded"},
	}

	for idx, pattern := range patterns {
		actual := utils.GetContentType(pattern.key)
		if pattern.expected != actual {
			t.Errorf("testcase%d: expected: %s, actual %s", idx, pattern.expected, actual)
		}
	}
}

func TestReadWriteFile(t *testing.T) {
	type testJson struct {
		Test string
	}
	testText := "Hello Redmine World"
	filedir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(filedir)
	filename := filepath.Join(filedir, "dummy.json")
	testByte, _ := json.Marshal(testJson{testText})
	err = utils.WriteFile(filename, testByte)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	var actualJson testJson
	body, err := utils.ReadFile(filename)
	_ = json.Unmarshal(body, &actualJson)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	if actualJson.Test != testText {
		t.Errorf("expected: %s, actual %s", testText, actualJson.Test)
	}
}

func TestWriteFileExistDir(t *testing.T) {
	fileDir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(fileDir)
	fileDir2 := filepath.Join(fileDir, "/dummy")
	filename := filepath.Join(fileDir, "/dummy")
	err = utils.MakeDir(fileDir2)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	err = utils.WriteFile(filename, []byte("aaaa"))
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestReadFileInvalidFileName(t *testing.T) {
	fileDir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(fileDir)
	filename := filepath.Join(fileDir, "/dummy")
	err = utils.MakeDir(fileDir)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	_, err = utils.ReadFile(filename)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestReadFileLookUpDir(t *testing.T) {
	fileDir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(fileDir)
	fileDir2 := filepath.Join(fileDir, "/dummy")
	filename := filepath.Join(fileDir, "/dummy")
	err = utils.MakeDir(fileDir2)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	_, err = utils.ReadFile(filename)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestCheckDir(t *testing.T) {
	fileDir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(fileDir)
	actual := utils.CheckDir(fileDir)
	if !actual {
		t.Errorf("expected: %t, actual %t", true, actual)
	}
}

func TestCheckDirNoDir(t *testing.T) {
	fileDir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(fileDir)
	fileDir = filepath.Join(fileDir, "/test")
	actual := utils.CheckDir(fileDir)
	if actual {
		t.Errorf("expected: %t, actual %t", false, actual)
	}
}

func TestMakeDir(t *testing.T) {
	fileDir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(fileDir)
	fileDir = filepath.Join(fileDir, "/test/test")
	err = utils.MakeDir(fileDir)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
}

func TestMakeDirExistsDir(t *testing.T) {
	fileDir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(fileDir)
	fileDir = filepath.Join(fileDir, "/test/test")
	err = utils.MakeDir(fileDir)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	err = utils.MakeDir(fileDir)
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
}

func TestMakeDirExistsFile(t *testing.T) {
	fileDir, err := ioutil.TempDir("", "redump_test")
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	defer os.RemoveAll(fileDir)
	fileDir2 := filepath.Join(fileDir, "/dummy")
	filename := filepath.Join(fileDir, "/dummy")
	err = utils.WriteFile(filename, []byte("aaaa"))
	if err != nil {
		t.Errorf("Error occurred: %s", err)
	}
	err = utils.MakeDir(fileDir2)
	if err == nil {
		t.Errorf("Error not occurred")
	}
}

func TestSanitizeInvalidFileName(t *testing.T) {
	testStr := "test test　testtest.com"
	actual := utils.SanitizeInvalidFileName(testStr)
	if actual != "test_test_testtest.com" {
		t.Errorf("expected: %s, actual %s", "test_test_testtest.com", testStr)
	}
}

func ExampleCheckDir() {
	resp := utils.CheckDir("test")
	fmt.Printf("%t", resp)
}

func ExampleMakeDir() {
	_ = utils.MakeDir("test")
	resp := utils.CheckDir("test")
	// true
	fmt.Printf("%t", resp)
}

func ExampleGetContentType() {
	resp := utils.GetContentType("test.png")
	// image/png
	fmt.Println(resp)
}

func ExampleReadFile() {
	resp, _ := utils.ReadFile("test.json")
	fmt.Println(string(resp))
}

func ExampleWriteFile() {
	_ = utils.WriteFile("test.json", []byte("{\"some_json\": \"tubone24\"}"))
}

func ExampleSanitizeInvalidFileName() {
	resp := utils.SanitizeInvalidFileName("test test　testtest test 　test")
	// testtesttesttesttesttest
	fmt.Println(resp)
}
