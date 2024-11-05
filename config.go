package main

import (
	"encoding/json"
	"errors"
	"os"
)

const (
	DefaultProxyServerPort = "8888"
	DefaultWebServerPort   = "8080"
)

type Config struct {
	ProxyServerPort string `json:"proxyServerPort"`
	WebServerPort   string `json:"webServerPort"`
	CertPath        string `json:"certPath"`
	KeyPath         string `json:"keyPath"`
	RuleDbPath      string `json:"ruleDbPath"`
	RequestDbPath   string `json:"requestDbPath"`
	WebServerPath   string `json:"webServerPath"`
}

func getConfig() Config {

	certPath, keyPath := getCertKeyPath()
	ruleDbPath, requestDbPath := getDatabasePaths()
	config := Config{
		ProxyServerPort: DefaultProxyServerPort,
		WebServerPort:   DefaultWebServerPort,
		RuleDbPath:      ruleDbPath,
		RequestDbPath:   requestDbPath,
		CertPath:        certPath,
		KeyPath:         keyPath,
	}

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

func saveConfig(c Config) {
	configPath := getConfigPath()
	configFile, err := os.Create(configPath)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	jsonParser := json.NewEncoder(configFile)
	jsonParser.Encode(c)
}
