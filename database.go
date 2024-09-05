package main

type RedirectRecord struct {
	If   string
	Then string
}

type Record struct {
	redirect RedirectRecord
}

type Collection struct {
	records []Record
}

type Database interface {
	// InsertCollection(record Record) error
	GetCollection() (Collection, error)
	GetCollections() ([]Collection, error)
}

type FileDatabase struct {
	filePath string
}

func (f *FileDatabase) Insert(record RedirectRecord) error {
	return nil
}

func (f *FileDatabase) FindAll() ([]RedirectRecord, error) {
	return []RedirectRecord{}, nil
}

func (f *FileDatabase) FindOne(ifValue string) (RedirectRecord, error) {
	return RedirectRecord{}, nil
}
