package main

import (
	"testing"
)

func TestFileDatabase(t *testing.T) {
	test_database := "database_test.json"

	fileDatabase := FileDatabase{filePath: test_database}
	fileDatabase.load()

	expected := []Collection{
		{Name: "collection 1", Records: []Record{{Type: "redirect", If: "url contains \"example\"", Then: "example.com"}}},
		{Name: "collection 1"},
	}

	if len(expected) != len(fileDatabase.collections) {
		t.Error("Exprected length", len(expected), ", but got", len(fileDatabase.collections))
	}

	for i := range expected {
		collectionName := expected[i].Name
		dbCollection, err := fileDatabase.GetCollection(collectionName)
		if err != nil {
			t.Error("Collection", collectionName, "not found")
		}

		if collectionName != dbCollection.Name {
			t.Error("Expected name", collectionName)
		}

		for j := range expected[i].Records {
			if expected[i].Records[j] != dbCollection.Records[j] {
				t.Error("Record not matching")
			}
		}
	}
}
