package config

import (
	"bytes"
	"github.com/umirode/prot/tools"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//go:generate go-bindata -o prot.go -pkg config templates
type Config struct {
	Lang    string            `yaml:"Lang"`
	Modules map[string]Module `yaml:"Modules"`
}

type Module struct {
	Repository string `yaml:"Repository"`
	Auth       *struct {
		Type   string            `yaml:"Type"`
		Config map[string]string `yaml:"Config"`
	} `yaml:"Auth"`
}

func NewConfig(filePath string) (*Config, error) {
	configFileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	schemaJsonAsset, err := templatesSchemaJson()
	if err != nil {
		return nil, err
	}

	err = tools.ValidateYaml(configFileData, bytes.NewReader(schemaJsonAsset.bytes))
	if err != nil {
		//return nil, err
	}

	config := &Config{}

	err = yaml.Unmarshal(configFileData, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func GetConfigTemplate() (string, error) {
	protYamlAsset, err := templatesProtYaml()
	if err != nil {
		return "nil", err
	}

	return string(protYamlAsset.bytes), nil
}
