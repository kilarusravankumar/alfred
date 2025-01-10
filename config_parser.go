package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	BuildName    string `json:"build_name"`
	BuildVersion string `json:"build_version"`
}

func ParseConfig() Config {
	fileName := "alfred_config.json"
	file, err := os.Open(fileName)
	if err != nil {
		logger.Error(err.Error())
	}
	jsonDecoder := json.NewDecoder(file)
	config := Config{}
	err = jsonDecoder.Decode(&config)
	if err != nil {
		logger.Error(err.Error())

	}
	return config

}
