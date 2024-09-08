package main

import (
	"context"
	"fmt"
)

type App struct {
	ctx      context.Context
	proxy    *Proxy
	database Database
}

type ReturnValue struct {
	Collections []Collection `json:"collections"`
	Collection  Collection   `json:"collection"`
	Error       string       `json:"error"`
}

func NewApp() *App {

	database := FileDatabase{
		filePath: "database.json",
	}
	database.load()

	return &App{
		proxy:    NewProxy("example.crt", "example.key"),
		database: &database,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetCollections() []Collection {
	return a.database.GetCollections()
}

func (a *App) GetCollection(collectionId int) ReturnValue {
	collection, err := a.database.GetCollection(collectionId)

	if err != nil {
		return ReturnValue{
			Error: err.Error(),
		}
	}

	return ReturnValue{
		Collection: *collection,
	}
}

func (a *App) NewRecord(collectionId int, recordType string) {
	a.database.Insert(collectionId, Record{Type: recordType})

}

func (a *App) SaveRecord(collectionId int, recordId int, updatedRecord Record) {
	a.database.Update(collectionId, recordId, updatedRecord)

}

func (a *App) DeleteRecord(collectionId int, recordId int) {
	a.database.DeleteRecord(collectionId, recordId)
}
