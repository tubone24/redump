package config_test

import (
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
	filename := filepath.FromSlash(dir + "/../../config.toml.example")
	cfg, err := config.GetConfig(filename)
	if err != nil {
		t.Errorf("Error occured %s", err)
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
