package main

import (
	"os"
	"path/filepath"
)

const (
	AppName  = "middleman"
	CertName = "ca.crt"
	KeyName  = "ca.key"
)

func getAppConfigDir() string {
	dirName, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	appConfigDir := filepath.Join(dirName, AppName)
	return appConfigDir
}

func getCertKeyPath() (string, string) {
	appConfigDir := getAppConfigDir()
	certPath, keyPath := filepath.Join(appConfigDir, CertName), filepath.Join(appConfigDir, KeyName)
	os.MkdirAll(filepath.Dir(certPath), os.ModePerm)
	os.MkdirAll(filepath.Dir(keyPath), os.ModePerm)
	return certPath, keyPath
}

func getConfigPath() string {
	appConfigDir := getAppConfigDir()
	configPath := filepath.Join(appConfigDir, "config.json")
	os.MkdirAll(filepath.Dir(configPath), os.ModePerm)
	return configPath
}
