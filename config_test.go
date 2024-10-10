package main

import (
	"fmt"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	testConfigDir := "tmp/middlemantest"
	createDirsOrPanic(testConfigDir)

	getAppConfigDir = func() string {
		return testConfigDir
	}

	// Test when config.json is not there
	os.Remove(fmt.Sprintf("%s/config.json", testConfigDir))
	config := getConfig()
	if config.ServerPort != "8888" {
		t.Error("Expected 8888, got ", config.ServerPort)
	}
	if config.CertPath != fmt.Sprintf("%s/ca.crt", testConfigDir) {
		t.Error("Expected ", fmt.Sprintf("%s/ca.crt", testConfigDir), " got ", config.CertPath)
	}
	if config.KeyPath != fmt.Sprintf("%s/ca.key", testConfigDir) {
		t.Error("Expected ", fmt.Sprintf("%s/ca.key", testConfigDir), " got ", config.KeyPath)
	}
	if config.DatabasePath != fmt.Sprintf("%s/database.json", testConfigDir) {
		t.Error("Expected ", fmt.Sprintf("%s/database.json", testConfigDir), " got ", config.DatabasePath)
	}

	// Test when config.json is there
	saveConfig(Config{
		ServerPort:   "1111",
		CertPath:     "cert",
		KeyPath:      "key",
		DatabasePath: "db",
	})
	config = getConfig()
	if config.ServerPort != "1111" {
		t.Error("Expected 1111, got ", config.ServerPort)
	}
	if config.CertPath != "cert" {
		t.Error("Expected cert, got ", config.CertPath)
	}
	if config.KeyPath != "key" {
		t.Error("Expected key, got ", config.KeyPath)
	}
	if config.DatabasePath != "db" {
		t.Error("Expected db, got ", config.DatabasePath)
	}

	// When config is invalid
	os.Remove(fmt.Sprintf("%s/config.json", testConfigDir))
	configFile, err := os.Create(fmt.Sprintf("%s/config.json", testConfigDir))
	if err != nil {
		t.Error(err)
	}
	configFile.Write([]byte("foo"))
	configFile.Close()
	config = getConfig()
	if config.ServerPort != "8888" {
		t.Error("Expected 8888, got ", config.ServerPort)
	}
	if config.CertPath != fmt.Sprintf("%s/ca.crt", testConfigDir) {
		t.Error("Expected ", fmt.Sprintf("%s/ca.crt", testConfigDir), " got ", config.CertPath)
	}
	if config.KeyPath != fmt.Sprintf("%s/ca.key", testConfigDir) {
		t.Error("Expected ", fmt.Sprintf("%s/ca.key", testConfigDir), " got ", config.KeyPath)
	}
	if config.DatabasePath != fmt.Sprintf("%s/database.json", testConfigDir) {
		t.Error("Expected ", fmt.Sprintf("%s/database.json", testConfigDir), " got ", config.DatabasePath)
	}

}
