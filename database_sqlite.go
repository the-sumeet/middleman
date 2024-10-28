package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

const (
	// Tables
	MIGRATION = "migration"
	RULE      = "rule"

	// Error contains
	ERR_NO_SUCH_TABLE     = "no such table"
	ERR_FILE_NOT_DATABASE = "file is not a database"
)

type SqliteDatabase struct {
	filePath string
	db       *sql.DB
}

func (s *SqliteDatabase) migrateVersion(from int) {
	for i := from; i < len(UP); i++ {
		ddl := UP[i]
		_, err := s.db.Exec(ddl)
		if err != nil {
			panic(err)
		}
	}
}

func (s *SqliteDatabase) migrate() {
	row := s.db.QueryRow(fmt.Sprintf("SELECT MAX(version) FROM %s", MIGRATION))

	if row.Err() != nil {
		if strings.Contains(row.Err().Error(), ERR_NO_SUCH_TABLE) {
			s.migrateVersion(0)
		} else if strings.Contains(row.Err().Error(), ERR_FILE_NOT_DATABASE) {
			// Panic
		}
	} else {
		var version int
		row.Scan(&version)
		s.migrateVersion(version + 1)
	}

}

func NewSqliteDatabase(filePath string) *SqliteDatabase {
	db, err := sql.Open("sqlite", filePath)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	database := &SqliteDatabase{db: db, filePath: filePath}
	database.migrate()
	return database
}

func (s *SqliteDatabase) AddRule(rule Rule) (any, error) {
	recordId := time.Now().UTC().Unix()
	var res sql.Result

	marshalled, err := json.Marshal(rule)
	if err != nil {
		return recordId, err
	}

	res, err = s.db.Exec(fmt.Sprintf("INSERT INTO %s (id, data) VALUES (?, ?)", RULE), recordId, marshalled)
	if err != nil {
		return nil, err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return lastInsertId, nil
}

func (s *SqliteDatabase) GetOneRule(id any) (Rule, error) {

	var jsonData string
	var rule Rule

	row := s.db.QueryRow(fmt.Sprintf("SELECT id, data FROM %s WHERE id = ?", RULE), id)

	err := row.Scan(&id, &jsonData)
	if err != nil {
		return rule, err
	}

	err = json.Unmarshal([]byte(jsonData), &rule)
	if err != nil {
		return rule, err
	}

	return rule, err
}

func (s *SqliteDatabase) GetManyRules(recordType string) ([]Rule, error) {
	rules := []Rule{}

	rows, err := s.db.Query(fmt.Sprintf("SELECT id, data FROM %s WHERE json_extract(data, '$.type') = ?", RULE), recordType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rule Rule
		var jsonData string
		var id int

		err := rows.Scan(&id, &jsonData)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(jsonData), &rule)
		if err != nil {
			return nil, err
		}

		rules = append(rules, rule)
	}

	return rules, nil
}
