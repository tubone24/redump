package config

import "github.com/BurntSushi/toml"

type ServerConfig struct {
	Url       string `toml:"url"`
	Key       string `toml:"key"`
	ProjectId int    `toml:"project_id"`
	Sleep     int    `toml:"sleep"`
	Timeout   int    `toml:"timeout"`
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
	return &config, nil
}
