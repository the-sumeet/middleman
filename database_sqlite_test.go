package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

const SqlFile = "/tmp/middlemantest/database.db"

func TestSqliteDatabase(t *testing.T) {

	db := NewSqliteDatabase(SqlFile)
	fmt.Println(db)
}

func TestSqliteCRUD(t *testing.T) {

	os.Remove(SqlFile)
	db := NewSqliteDatabase(SqlFile)

	// Add and Get a rule
	rule := Rule{
		Type:    REDIRECT,
		Enabled: true,
		Entity:  "url",
		Op:      "contains",
		Value:   "foo.com",
	}
	insertedId, _ := db.AddRule(rule)
	dBRule, err := db.GetOneRule(insertedId)
	if err != nil {
		t.Error(err)
	}
	if dBRule.Enabled != true || dBRule.Entity != "url" || dBRule.Op != "contains" || dBRule.Value != "foo.com" {
		t.Errorf("Expected %v, got %v", rule, dBRule)
	}

	// Try to get with invalid id
	expected := errors.New("sql: no rows in result set")
	_, err = db.GetOneRule(999)
	if err.Error() != expected.Error() {
		t.Errorf("Expected %v, got %v", expected, err)
	}

	// Get many
	records, _ := db.GetManyRules(REDIRECT)
	if len(records) != 1 {
		t.Errorf("Expected 1, got %d", len(records))
	}

}
