package config

import (
	"github.com/BurntSushi/toml"
)

type MissingConfigError struct {
	ConfigName string
}

func (e *MissingConfigError) Error() string {
	return "Missing Config: " + e.ConfigName + " is Required"
}

type ServerConfig struct {
	Url       string `toml:"url"`
	Key       string `toml:"key"`
	ProjectId int    `toml:"project_id"`
	Sleep     int    `toml:"sleep"`
	Timeout   int    `toml:"timeout"`
	ProxyUrl   string    `toml:"proxy_url"`
}

type MappingValue struct {
	Old int `toml:"old"`
	New int `toml:"new"`
}

type Mapping struct {
	Name    string         `toml:"name"`
	Default int            `toml:"default"`
	Values  []MappingValue `toml:"values"`
}

type Config struct {
	ServerConfig    ServerConfig `toml:"server"`
	NewServerConfig ServerConfig `toml:"new_server"`
	Mappings        []Mapping    `toml:"mappings"`
}

// GetConfig is a function to read the configuration described in toml.
// Please refer to README for how to write config.toml.
// https://github.com/tubone24/redump#precondition
func GetConfig(configPath string) (*Config, error) {
	var config Config
	if configPath == "" {
		_, err := toml.DecodeFile("config.toml", &config)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := toml.DecodeFile(configPath, &config)
		if err != nil {
			return nil, err
		}
	}
	if config.ServerConfig.Url == "" {
		return nil, &MissingConfigError{ConfigName: "server.url"}
	}
	if config.NewServerConfig.Url == "" {
		return nil, &MissingConfigError{ConfigName: "new_server.url"}
	}
	if config.ServerConfig.Timeout == 0 {
		config.ServerConfig.Timeout = 60000
	}
	if config.NewServerConfig.Timeout == 0 {
		config.NewServerConfig.Timeout = 60000
	}
	return &config, nil
}
