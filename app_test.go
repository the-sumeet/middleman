package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

const (
	testConfigDir = "/tmp/middlemantest"
)

func setup() {
	os.RemoveAll("/tmp/middlemantest")
	createDirsOrPanic(testConfigDir)
	getAppConfigDir = func() string {
		return testConfigDir
	}
}

func TestAppConfig(t *testing.T) {

	setup()

	app := NewApp()

	// Check config
	if app.config.ServerPort != "8888" {
		t.Error("Expected 8888, got ", app.config.ServerPort)
	}
	if app.config.CertPath != fmt.Sprintf("%s/ca.crt", testConfigDir) {
		t.Error("Expected ", fmt.Sprintf("%s/ca.crt", testConfigDir), " got ", app.config.CertPath)
	}
	if app.config.KeyPath != fmt.Sprintf("%s/ca.key", testConfigDir) {
		t.Error("Expected ", fmt.Sprintf("%s/ca.key", testConfigDir), " got ", app.config.KeyPath)
	}
	if app.config.DatabasePath != fmt.Sprintf("%s/database.json", testConfigDir) {
		t.Error("Expected ", fmt.Sprintf("%s/database.json", testConfigDir), " got ", app.config.DatabasePath)
	}

}

func TestGetConfig(t *testing.T) {

	setup()

	app := NewApp()

	if app.config != app.GetConfig() {
		t.Error("Expected ", app.config, " got ", app.GetConfig())
	}

}

func TestAddConfigPort(t *testing.T) {

	setup()

	app := NewApp()

	app.AddConfigPort("9999")
	if app.config.ServerPort != "9999" {
		t.Error("Expected 9999, got ", app.config.ServerPort)
	}

	// Read config file
	configFromFile := getConfig()
	if configFromFile.ServerPort != "9999" {
		t.Error("Expected 9999, got ", configFromFile.ServerPort)
	}

}

func TestGetMany(t *testing.T) {

	setup()

	// Test with empty json file
	data, err := json.Marshal(map[string]any{})
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(getDatabasePath(), data, 0644)
	if err != nil {
		panic(err)
	}
	app := NewApp()
	res := app.GetMany(REDIRECT)
	if len(res.Cancels) != 0 ||
		len(res.Delays) != 0 ||
		len(res.ModifyHeaders) != 0 ||
		len(res.ModifyRequestBody) != 0 ||
		len(res.ModifyResponseBody) != 0 ||
		len(res.Redirects) != 0 {
		t.Error("Expected empty array, got ", res)
	}

	// Test with databse file with values
	redirects := []Redirect{
		{Request: Request{Enabled: true, Entity: "url", Op: "equal", Value: "example"}, ToValue: "example2"},
		{Request: Request{Enabled: true, Entity: "host", Op: "contains", Value: "foo"}, ToValue: "bar"},
	}
	cancels := []Cancel{
		{Enabled: false, Entity: "url", Op: "equal", Value: "example"},
	}
	delays := []Delay{
		{Request: Request{Enabled: true, Entity: "url", Op: "equal", Value: "example"}, DelaySec: 10},
	}
	modifyHeaders := []ModifyHeader{
		{Request: Request{Enabled: true, Entity: "url", Op: "equal", Value: "example"}, Mods: []Header{{Action: "add", IsRequest: true, Name: "bar", Value: "baz"}}},
	}
	modifyRequestBody := []ModifyRequestBody{
		{Request: Request{Enabled: true, Entity: "url", Op: "equal", Value: "example"}, Body: "foo"},
	}
	modifyResponseBody := []ModifyResponseBody{
		{Request: Request{Enabled: true, Entity: "url", Op: "equal", Value: "example"}, Body: "foo"},
	}

	data, err = json.Marshal(map[string]any{
		"redirects":          redirects,
		"cancels":            cancels,
		"delays":             delays,
		"modifyHeaders":      modifyHeaders,
		"modifyRequestBody":  modifyRequestBody,
		"modifyResponseBody": modifyResponseBody,
	})
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(getDatabasePath(), data, 0644)
	if err != nil {
		panic(err)
	}
	app = NewApp()
	res = app.GetMany(REDIRECT)
	if len(res.Redirects) != 2 {
		t.Error("Expected 2 redirects, got ", len(res.Redirects))
	}
	if res.Redirects[0] != redirects[0] {
		t.Error("Expected ", redirects[0], " got ", res.Redirects[0])
	}
	if res.Redirects[1] != redirects[1] {
		t.Error("Expected ", redirects[1], " got ", res.Redirects[1])
	}
	res = app.GetMany(CANCEL)
	if len(res.Cancels) != 1 {
		t.Error("Expected 1 cancel, got ", len(res.Cancels))
	}
	if res.Cancels[0] != cancels[0] {
		t.Error("Expected ", cancels[0], " got ", res.Cancels[0])
	}
	res = app.GetMany(DELAY)
	if len(res.Delays) != 1 {
		t.Error("Expected 1 delay, got ", len(res.Delays))
	}
	if res.Delays[0] != delays[0] {
		t.Error("Expected ", delays[0], " got ", res.Delays[0])
	}
	res = app.GetMany(MODIFY_HEADERS)
	if len(res.ModifyHeaders) != 1 {
		t.Error("Expected 1 modifyHeader, got ", len(res.ModifyHeaders))
	}
	if res.ModifyHeaders[0].Enabled != modifyHeaders[0].Enabled ||
		res.ModifyHeaders[0].Entity != modifyHeaders[0].Entity ||
		res.ModifyHeaders[0].Op != modifyHeaders[0].Op ||
		res.ModifyHeaders[0].Value != modifyHeaders[0].Value ||
		len(res.ModifyHeaders[0].Mods) != 1 || res.ModifyHeaders[0].Mods[0] != modifyHeaders[0].Mods[0] {
		t.Error("Expected ", modifyHeaders[0], " got ", res.ModifyHeaders[0])
	}
	res = app.GetMany(MODIFY_REQUEST_BODY)
	if len(res.ModifyRequestBody) != 1 {
		t.Error("Expected 1 modifyRequestBody, got ", len(res.ModifyRequestBody))
	}
	if res.ModifyRequestBody[0] != modifyRequestBody[0] {
		t.Error("Expected ", modifyRequestBody[0], " got ", res.ModifyRequestBody[0])
	}
	res = app.GetMany(MODIFY_RESPONSE_BODY)
	if len(res.ModifyResponseBody) != 1 {
		t.Error("Expected 1 modifyResponseBody, got ", len(res.ModifyResponseBody))
	}

}

func TestSave(t *testing.T) {

	setup()
	app := NewApp()

	app.Add(
		REDIRECT,
		InValue{Redirect: Redirect{Request: Request{Enabled: true, Entity: "url", Op: "equal", Value: "example"}, ToValue: "example2"}},
	)

}
