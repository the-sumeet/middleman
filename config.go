package main

import (
	"encoding/json"
	"errors"
	"os"
)

const (
	DefaultServerPort = "8888"
)

type Config struct {
	// Database struct {
	// 	Host     string `json:"host"`
	// 	Password string `json:"password"`
	// 	User     string `json:"user"`
	// 	Port     int    `json:"port"`
	// } `json:"database"`
	ServerPort string `json:"serverPort"`
}

func getConfig() Config {

	config := Config{ServerPort: DefaultServerPort}

	configPath := getConfigPath()

	configFile, err := os.Open(configPath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}

	} else {
		defer configFile.Close()

		jsonParser := json.NewDecoder(configFile)
		jsonParser.Decode(&config)
	}
	return config
}
