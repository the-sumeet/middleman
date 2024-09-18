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
	ServerPort string `json:"serverPort"`
	CertPath   string `json:"certPath"`
	KeyPath    string `json:"keyPath"`
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

	certPath, keyPath := getCertKeyPath()
	config.CertPath = certPath
	config.KeyPath = keyPath
	return config
}
