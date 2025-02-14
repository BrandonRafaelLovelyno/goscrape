package cli

import (
	"encoding/json"
	"fmt"

	_json "github.com/BrandonRafaelLovelyno/goscrape/pkg/json"
)

func attachConfig(confDir *string, cmdInput *CommandInput) error {
	if confDir == nil {
		return fmt.Errorf("no configuration file directory provided")
	}

	config, err := extractConfigJson(*confDir)
	if err != nil {
		return fmt.Errorf("failed to extract config json: %v", err.Error())
	}

	cmdInput.TargetSelectors = config.targetSelectors
	cmdInput.WaitedSelectors = config.waitedSelectors

	return nil
}

func unmarshalConfigJson(jsonData []byte) (*Config, error) {
	var config Config

	err := json.Unmarshal(jsonData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func extractConfigJson(confDir string) (*Config, error) {
	file, err := _json.GetJsonData(confDir)
	if err != nil {
		return nil, fmt.Errorf("failed to get json data: %v", err.Error())
	}

	config, err := unmarshalConfigJson(*file)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json data: %v", err.Error())
	}

	return config, nil
}
