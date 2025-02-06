package cli

import (
	"encoding/json"
	pkgj "github.com/BrandonRafaelLovelyno/goscrape/pkg/json"
)

func unmarshalConfigJson(jsonData []byte) (*Config, error) {
	var config Config

	err := json.Unmarshal(jsonData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func extractConfigJson(confDir string) (*Config, error) {
	file, err := pkgj.GetJsonData(confDir)
	if err != nil {
		return nil, err
	}

	config, err := unmarshalConfigJson(*file)
	if err != nil {
		return nil, err
	}

	return config, nil
}
