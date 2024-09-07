package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type RedirectRecord struct {
	If   string
	Then string
}

type Record struct {
	Type string `json:"type"`
	If   string `json:"if"`
	Then string `json:"then"`
}

type Collection struct {
	Name    string   `json:"name"`
	Records []Record `json:"records"`
}

type Database interface {
	// InsertCollection(record Record) error
	// GetCollection() (Collection, error)
	GetCollections() []Collection
	GetCollection(int) (Collection, error)
	Update(collectionId int, recordId int, record Record) error
}

type FileDatabase struct {
	filePath    string
	collections []Collection
}

func (f *FileDatabase) load() {
	data, err := os.ReadFile(f.filePath)
	if err != nil {
		panic(err)
	}

	var database struct {
		Collections []Collection `json:"collections"`
	}
	err = json.Unmarshal(data, &database)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	f.collections = database.Collections

}

func (f *FileDatabase) store() {
	data, err := json.Marshal(map[string]any{
		"collections": f.collections,
	})
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(f.filePath, data, 0644)
	if err != nil {
		panic(err)
	}
}

func (f *FileDatabase) GetCollections() []Collection {
	return f.collections
}

func (f *FileDatabase) GetCollection(collectionId int) (Collection, error) {
	// Index based id
	if collectionId >= len(f.collections) {
		return Collection{}, fmt.Errorf("collection with id %s not found", collectionId)
	}

	return f.collections[collectionId], nil
}

func (f *FileDatabase) Update(collectionId int, recordId int, record Record) error {

	collection, err := f.GetCollection(collectionId)
	if err != nil {
		return err
	}

	if recordId >= len(collection.Records) {
		return fmt.Errorf("record %s not found", recordId)
	}

	collection.Records[recordId] = record
	f.store()
	return nil
}
