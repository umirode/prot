package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Lang    string            `yaml:"Lang"`
	Modules map[string]Module `yaml:"Modules"`
}

type Module struct {
	Url        string             `yaml:"Url"`
	AuthType   *string            `yaml:"AuthType"`
	AuthConfig *map[string]string `yaml:"AuthConfig"`
}

func NewConfig(filePath string) (*Config, error) {
	configFileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = yaml.Unmarshal(configFileData, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
