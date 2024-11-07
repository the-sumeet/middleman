package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	_ "modernc.org/sqlite"
)

const (
	// Tables
	MIGRATION = "migration"
	RULE      = "rule"
	REQUEST   = "request"

	// Error contains
	ERR_NO_SUCH_TABLE     = "no such table"
	ERR_FILE_NOT_DATABASE = "file is not a database"
)

type SqliteDatabase struct {
	filePath  string
	ruleDb    *sql.DB
	requestDb *sql.DB
}

func (s *SqliteDatabase) migrateVersion(from int, db *sql.DB) {
	for i := from; i < len(UP); i++ {
		ddl := UP[i]
		_, err := db.Exec(ddl)
		if err != nil {
			panic(err)
		}
	}
}

func (s *SqliteDatabase) migrate() {
	// Migrate rule database
	row := s.ruleDb.QueryRow(fmt.Sprintf("SELECT MAX(version) FROM %s", MIGRATION))
	if row.Err() != nil {
		if strings.Contains(row.Err().Error(), ERR_NO_SUCH_TABLE) {
			s.migrateVersion(0, s.ruleDb)
		} else if strings.Contains(row.Err().Error(), ERR_FILE_NOT_DATABASE) {
			// Panic
		}
	} else {
		var version int
		row.Scan(&version)
		s.migrateVersion(version+1, s.ruleDb)
	}

	// Migrate requests database
	row = s.requestDb.QueryRow(fmt.Sprintf("SELECT MAX(version) FROM %s", MIGRATION))
	if row.Err() != nil {
		if strings.Contains(row.Err().Error(), ERR_NO_SUCH_TABLE) {
			s.migrateVersion(0, s.requestDb)
		} else if strings.Contains(row.Err().Error(), ERR_FILE_NOT_DATABASE) {
			// Panic
		}
	} else {
		var version int
		row.Scan(&version)
		s.migrateVersion(version+1, s.requestDb)
	}
}

func NewSqliteDatabase(ruleDbPath, requestDbPath string) *SqliteDatabase {
	// Open databae, set journal mode to WAL, close and reopen.
	ruleDb, err := sql.Open("sqlite", ruleDbPath)
	if err != nil {
		panic(err)
	}
	_, err = ruleDb.Exec("PRAGMA journal_mode=WAL;")
	ruleDb.Close()
	ruleDb, err = sql.Open("sqlite", ruleDbPath)
	if err != nil {
		panic(err)
	}

	// Open databae, set journal mode to WAL, close and reopen.
	requestDb, err := sql.Open("sqlite", requestDbPath)
	if err != nil {
		panic(err)
	}
	_, err = requestDb.Exec("PRAGMA journal_mode=WAL;")
	requestDb.Close()
	requestDb, err = sql.Open("sqlite", requestDbPath)
	requestDb.Exec("DELETE FROM request") // Clear table on start.
	database := &SqliteDatabase{ruleDb: ruleDb, requestDb: requestDb, filePath: ruleDbPath}
	database.migrate()

	return database
}

func (s *SqliteDatabase) AddRule(rule Rule) (any, error) {
	var res sql.Result

	marshalled, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}

	res, err = s.ruleDb.Exec(fmt.Sprintf("INSERT INTO %s (data) VALUES (?)", RULE), marshalled)
	if err != nil {
		return nil, err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return lastInsertId, nil
}

func (s *SqliteDatabase) UpdateRule(id any, rule Rule) (Rule, error) {

	marshalled, err := json.Marshal(rule)
	if err != nil {
		return rule, err
	}

	_, err = s.ruleDb.Exec(fmt.Sprintf("UPDATE %s SET data = ? WHERE id = ?", RULE), marshalled, id)
	if err != nil {
		return rule, err
	}

	updatedRule, err := s.GetOneRule(id)
	if err != nil {
		return rule, err
	}

	return updatedRule, nil
}

func (s *SqliteDatabase) GetOneRule(id any) (Rule, error) {

	var jsonData string
	var rule Rule

	row := s.ruleDb.QueryRow(fmt.Sprintf("SELECT id, data FROM %s WHERE id = ?", RULE), id)

	err := row.Scan(&id, &jsonData)
	if err != nil {
		return rule, err
	}

	err = json.Unmarshal([]byte(jsonData), &rule)
	if err != nil {
		return rule, err
	}

	rule.Id = id
	return rule, err
}

func (s *SqliteDatabase) GetManyRules(recordType string) ([]Rule, error) {
	rules := []Rule{}

	rows, err := s.ruleDb.Query(fmt.Sprintf("SELECT id, data FROM %s WHERE json_extract(data, '$.type') = ?", RULE), recordType)
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

		rule.Id = id

		rules = append(rules, rule)
	}

	return rules, nil
}

func (s *SqliteDatabase) RemoveRule(id any) error {

	_, err := s.ruleDb.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = ?", RULE), id)

	return err
}

// Requests CRUD

func (s *SqliteDatabase) AddRequest(requestId any, request *HttpRequestLog) (any, error) {

	marshalled, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	_, err = s.requestDb.Exec(fmt.Sprintf("INSERT INTO %s (id, data) VALUES (?, ?)", REQUEST), requestId, marshalled)
	if err != nil {
		return nil, err
	}

	// Sqlite returning wrong last insert id, check this.
	// lastInsertId, err := res.LastInsertId()
	// if err != nil {
	// 	return nil, err
	// }

	return requestId, nil
}

func (s *SqliteDatabase) UpdateRequest(id string, request HttpRequestLog) (HttpRequestLog, error) {

	marshalled, err := json.Marshal(request)
	if err != nil {
		return request, err
	}

	_, err = s.requestDb.Exec(fmt.Sprintf("UPDATE %s SET data = ? WHERE id = ?", REDIRECT), marshalled, id)
	if err != nil {
		return request, err
	}

	updatedRequest, err := s.GetOneRequest(id)
	if err != nil {
		return request, err
	}

	return updatedRequest, nil
}

func (s *SqliteDatabase) GetOneRequest(id string) (HttpRequestLog, error) {

	var jsonData string
	var request HttpRequestLog

	row := s.requestDb.QueryRow(fmt.Sprintf("SELECT id, data FROM %s WHERE id = ?", REQUEST), id)

	err := row.Scan(&id, &jsonData)
	if err != nil {
		return request, err
	}

	err = json.Unmarshal([]byte(jsonData), &request)
	if err != nil {
		return request, err
	}

	request.Id = id
	return request, err
}

func (s *SqliteDatabase) GetManyRequests() ([]HttpRequestLog, error) {
	requests := []HttpRequestLog{}

	// rows, err := s.requestDb.Query(fmt.Sprintf("SELECT id, data FROM %s", REQUEST))
	rows, err := s.requestDb.Query(fmt.Sprintf("SELECT id, json_remove(data, '$.requestHeaders', '$.responseHeaders', '$.requestBody', '$.responseBody') AS data FROM %s ORDER BY id ASC LIMIT -1 OFFSET %d", REQUEST, skip))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var request HttpRequestLog
		var jsonData string
		var id string

		err := rows.Scan(&id, &jsonData)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(jsonData), &request)
		if err != nil {
			return nil, err
		}

		request.Id = id

		requests = append(requests, request)
	}

	return requests, nil
}

func (s *SqliteDatabase) RemoveRequest(id any) error {

	_, err := s.requestDb.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = ?", REQUEST), id)

	return err
}

func (s *SqliteDatabase) SetHeaders(id any, headers http.Header, isRequest bool) error {
	field := "requestHeaders"
	if !isRequest {
		field = "responseHeaders"
	}

	marshalled, err := json.Marshal(headers)
	if err != nil {
		return err
	}

	_, err = s.requestDb.Exec(fmt.Sprintf("UPDATE %s SET data = json_replace (data, '$.%s', ?) WHERE id = ?", REQUEST, field), marshalled, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SqliteDatabase) SetBody(id any, body string, isRequest bool) error {
	field := "requestBody"
	if !isRequest {
		field = "responseBody"
	}

	_, err := s.requestDb.Exec(fmt.Sprintf("UPDATE %s SET data = json_replace (data, '$.%s', ?) WHERE id = ?", REQUEST, field), body, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SqliteDatabase) SetRequestCancelled(id any, value bool) error {

	_, err := s.requestDb.Exec(fmt.Sprintf("UPDATE %s SET data = json_replace (data, '$.cancelled', ?) WHERE id = ?", REQUEST), value, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SqliteDatabase) SetRequestDelayed(id any, value int) error {

	_, err := s.requestDb.Exec(fmt.Sprintf("UPDATE %s SET data = json_replace (data, '$.delayed', ?) WHERE id = ?", REQUEST), value, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SqliteDatabase) SetRequestRedirected(id any, value bool) error {

	_, err := s.requestDb.Exec(fmt.Sprintf("UPDATE %s SET data = json_replace (data, '$.redirected', ?) WHERE id = ?", REQUEST), value, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SqliteDatabase) SetRequestHeaderModified(id any, value bool) error {

	_, err := s.requestDb.Exec(fmt.Sprintf("UPDATE %s SET data = json_replace (data, '$.requestHeaderModified', ?) WHERE id = ?", REQUEST), value, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SqliteDatabase) SetResponseHeaderModified(id any, value bool) error {

	_, err := s.requestDb.Exec(fmt.Sprintf("UPDATE %s SET data = json_replace (data, '$.responseHeaderModified', ?) WHERE id = ?", REQUEST), value, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SqliteDatabase) SetRequestBodyModified(id any, value bool) error {

	_, err := s.requestDb.Exec(fmt.Sprintf("UPDATE %s SET data = json_replace (data, '$.requestBodyModified', ?) WHERE id = ?", REQUEST), value, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *SqliteDatabase) SetResponseBodyModified(id any, value bool) error {

	_, err := s.requestDb.Exec(fmt.Sprintf("UPDATE %s SET data = json_replace (data, '$.responseBodyModified', ?) WHERE id = ?", REQUEST), value, id)

	if err != nil {
		return err
	}

	return nil
}
