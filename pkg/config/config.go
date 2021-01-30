package config

import "github.com/BurntSushi/toml"

type ServerConfig struct {
	Url string `toml:"url"`
	Key string `toml:"key"`
}

type Config struct {
	ServerConfig ServerConfig `toml:"server"`
}

func GetConfig() (*Config, error) {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
