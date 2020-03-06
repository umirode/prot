package tools

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/santhosh-tekuri/jsonschema"
	"gopkg.in/yaml.v2"
	"io"
)

func ValidateYaml(yamlData []byte, jsonSchema io.Reader) error {
	validator := newYamlValidator()

	return validator.ValidateYaml(yamlData, jsonSchema)
}

type yamlValidator struct {
}

func newYamlValidator() *yamlValidator {
	return &yamlValidator{}
}

func (y *yamlValidator) formatInvalidKeyError(keyPrefix string, key interface{}) error {
	var location string
	if keyPrefix == "" {
		location = "at top level"
	} else {
		location = fmt.Sprintf("in %s", keyPrefix)
	}
	return errors.Errorf("Non-string key %s: %#v", location, key)
}

func (y *yamlValidator) convertToStringKeysRecursive(value interface{}, keyPrefix string) (interface{}, error) {
	if mapping, ok := value.(map[interface{}]interface{}); ok {
		dict := make(map[string]interface{})
		for key, entry := range mapping {
			str, ok := key.(string)
			if !ok {
				return nil, y.formatInvalidKeyError(keyPrefix, key)
			}
			var newKeyPrefix string
			if keyPrefix == "" {
				newKeyPrefix = str
			} else {
				newKeyPrefix = fmt.Sprintf("%s.%s", keyPrefix, str)
			}
			convertedEntry, err := y.convertToStringKeysRecursive(entry, newKeyPrefix)
			if err != nil {
				return nil, err
			}
			dict[str] = convertedEntry
		}
		return dict, nil
	}
	if list, ok := value.([]interface{}); ok {
		var convertedList []interface{}
		for index, entry := range list {
			newKeyPrefix := fmt.Sprintf("%s[%d]", keyPrefix, index)
			convertedEntry, err := y.convertToStringKeysRecursive(entry, newKeyPrefix)
			if err != nil {
				return nil, err
			}
			convertedList = append(convertedList, convertedEntry)
		}
		return convertedList, nil
	}
	return value, nil
}

func (y *yamlValidator) ValidateYaml(yamlData []byte, jsonSchema io.Reader) error {
	var configText interface{}
	err := yaml.Unmarshal(yamlData, &configText)
	if err != nil {
		return err
	}

	configText, err = y.convertToStringKeysRecursive(configText, "")
	if err != nil {
		return err
	}

	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", jsonSchema); err != nil {
		return err
	}
	schema, err := compiler.Compile("schema.json")
	if err != nil {
		return err
	}
	if err := schema.ValidateInterface(configText); err != nil {
		return err
	}

	return nil
}
