package config_test

import (
	"fmt"
	"github.com/tubone24/redump/pkg/config"
	"os"
	"path/filepath"
	"testing"
)

func TestGetConfig(t *testing.T) {
	expected := config.Config{
		ServerConfig:    config.ServerConfig{Url: "https://example.com", Key: "xxxxx", ProjectId: 1, Sleep: 3000, Timeout: 10000},
		NewServerConfig: config.ServerConfig{Url: "https://blog.tubone-project24.xyz", Key: "xxxxx"}}
	dir, _ := os.Getwd()
	filename := filepath.FromSlash(dir + "/../../tests/test_assets/config_test.toml")
	cfg, err := config.GetConfig(filename)
	if err != nil {
		t.Errorf("Error occurred %s", err)
	}
	if cfg.ServerConfig.Url != expected.ServerConfig.Url {
		t.Errorf("expected '%s', actual '%s'", expected.ServerConfig.Url, cfg.ServerConfig.Url)
	}
	if cfg.ServerConfig.Key != expected.ServerConfig.Key {
		t.Errorf("expected '%s', actual '%s'", expected.ServerConfig.Key, cfg.ServerConfig.Key)
	}
	if cfg.ServerConfig.ProjectId != expected.ServerConfig.ProjectId {
		t.Errorf("expected '%d', actual '%d'", expected.ServerConfig.ProjectId, cfg.ServerConfig.ProjectId)
	}
	if cfg.ServerConfig.Sleep != expected.ServerConfig.Sleep {
		t.Errorf("expected '%d', actual '%d'", expected.ServerConfig.Sleep, cfg.ServerConfig.Sleep)
	}
	if cfg.ServerConfig.Timeout != expected.ServerConfig.Timeout {
		t.Errorf("expected '%d', actual '%d'", expected.ServerConfig.Timeout, cfg.ServerConfig.Timeout)
	}
	if cfg.NewServerConfig.Url != expected.NewServerConfig.Url {
		t.Errorf("expected '%s', actual '%s'", cfg.NewServerConfig.Url, expected.NewServerConfig.Url)
	}
	if cfg.NewServerConfig.Key != expected.NewServerConfig.Key {
		t.Errorf("expected '%s', actual '%s'", cfg.NewServerConfig.Key, expected.NewServerConfig.Key)
	}
}

func TestGetConfigMissingServerUrl(t *testing.T) {
	dir, _ := os.Getwd()
	filename := filepath.FromSlash(dir + "/../../tests/test_assets/config_test_missing_server_url.toml")
	_, err := config.GetConfig(filename)
	if err == nil {
		t.Error("Error does not occurred")
		return
	}
	if err.Error() != "Missing Config: server.url is Required" {
		t.Errorf("unexpected error message. expected '%s', actual '%s'", "Missing Config: server.url is Required", err.Error())
	}
}

func TestGetConfigMissingNewServerUrl(t *testing.T) {
	dir, _ := os.Getwd()
	filename := filepath.FromSlash(dir + "/../../tests/test_assets/config_test_missing_new_server_url.toml")
	_, err := config.GetConfig(filename)
	if err == nil {
		t.Error("Error does not occurred")
		return
	}
	if err.Error() != "Missing Config: new_server.url is Required" {
		t.Errorf("unexpected error message. expected '%s', actual '%s'", "Missing Config: new_server.url is Required", err.Error())
	}
}

func TestGetConfigMissingServerTimeout(t *testing.T) {
	dir, _ := os.Getwd()
	filename := filepath.FromSlash(dir + "/../../tests/test_assets/config_test_missing_server_timeout.toml")
	actual, err := config.GetConfig(filename)
	if err != nil {
		t.Errorf("Error occurred %s", err)
	}
	if actual.ServerConfig.Timeout != 60000 {
		t.Errorf("expected '%d', actual '%d'", 60000, actual.ServerConfig.Timeout)
	}
	if actual.NewServerConfig.Timeout != 60000 {
		t.Errorf("expected '%d', actual '%d'", 60000, actual.NewServerConfig.Timeout)
	}
}

func TestGetConfigDefaultConfigLocation(t *testing.T) {
	dir, _ := os.Getwd()
	filename := filepath.FromSlash(dir + "/../../config.toml")
	_, err := os.Stat(filename)
	if err != nil {
		_, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			t.Errorf("Error occurred %s", err)
		}
		_, err = config.GetConfig("")
		if err == nil {
			t.Errorf("Unexpected Error occurred")
			return
		}
	}
	_, _ = config.GetConfig("")
}

func ExampleGetConfig() {
	cfg, _ := config.GetConfig("")
	fmt.Println(cfg.ServerConfig.Url)
}
