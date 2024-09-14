package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Redirect struct {
	Entity  string `json:"entity"`
	Op      string `json:"op"`
	Value   string `json:"value"`
	ToType  string `json:"toType"`
	ToValue string `json:"toValue"`
	Enabled bool   `json:"enabled"`
}

type Database interface {
	// InsertCollection(record Record) error
	// GetCollection() (Collection, error)
	GetRedirects() []Redirect
	SaveRedirect(redirectId int, redirect Redirect) error
	RemoveRedirect(redirectId int) error
	AddRedirect(redirect Redirect) error
}

type FileDatabase struct {
	filePath  string
	redirects []Redirect
}

func (f *FileDatabase) load() {
	data, err := os.ReadFile(f.filePath)
	if err != nil {
		panic(err)
	}

	var database struct {
		Redirects []Redirect `json:"redirects"`
	}
	err = json.Unmarshal(data, &database)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	f.redirects = database.Redirects

}

func (f *FileDatabase) store() {
	data, err := json.Marshal(map[string]any{
		"redirects": f.redirects,
	})
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(f.filePath, data, 0644)
	if err != nil {
		panic(err)
	}
}

func (f *FileDatabase) GetRedirects() []Redirect {
	return f.redirects
}

func (f *FileDatabase) SaveRedirect(redirectId int, redirect Redirect) error {

	if redirectId >= len(f.redirects) {
		return fmt.Errorf("redirect with id %s not found", redirectId)
	}

	f.redirects[redirectId] = redirect
	f.store()
	return nil
}

func (f *FileDatabase) RemoveRedirect(redirectId int) error {

	if redirectId >= len(f.redirects) {
		return fmt.Errorf("redirect with id %s not found", redirectId)
	}

	f.redirects = append(f.redirects[:redirectId], f.redirects[redirectId+1:]...)

	f.store()
	return nil
}

func (f *FileDatabase) AddRedirect(redirect Redirect) error {
	f.redirects = append(f.redirects, redirect)
	f.store()
	return nil
}
