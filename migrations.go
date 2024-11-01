package main

const (
	UP_0 = `CREATE TABLE rule (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    data JSONB NOT NULL
);

CREATE TABLE request (
	id TEXT PRIMARY KEY,
	data JSONB NOT NULL
);

CREATE TABLE migration (version INTEGER PRIMARY KEY);

INSERT INTO migration (version) VALUES (0);`

	DOWN_0 = `DROP TABLE migration;
DROP TABLE rules;`
)

var (
	UP   = [...]string{UP_0}
	DOWN = [...]string{DOWN_0}
)
