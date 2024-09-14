package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// var db *sql.DB

// const ()

// func migrate() {
// 	maxVersion := 0

// 	files, err := os.ReadDir("migrations")
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, file := range files {
// 		version, err := strconv.Atoi(file.Name())
// 		if err != nil {
// 			panic(err)
// 		}

// 		if version > maxVersion {
// 			maxVersion = version
// 		}
// 	}

// 	migrationVersion := 1
// 	row := db.QueryRow("SELECT * FROM version LIMIT 1;", nil)
// 	err = row.Scan(&migrationVersion)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("No rows found")
// 		} else if strings.Contains(err.Error(), "no such table") {
// 			// No table found, start from migration version 1.
// 		} else {
// 			panic(err)
// 		}
// 		return
// 	}

// 	// Migrate from migrationVersion to maxVersion
// 	for migrationVersion <= maxVersion {
// 		_, err := os.ReadFile(fmt.Sprintf("migrations/%d", migrationVersion))
// 		if err != nil {
// 			panic(err)
// 		}

// 		// _, err = db.Exec(string(migrationSQL))
// 		if err != nil {
// 			panic(err)
// 		}
// 		log.Printf("Migrated to version %d", migrationVersion)

// 		migrationVersion++
// 	}

// }

// func init() {
// 	var err error
// 	db, err = sql.Open("sqlite", getSqlitePath())

// 	if err != nil {
// 		panic(err)
// 	}
// 	migrate()
// }
