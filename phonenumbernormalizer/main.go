package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./numbers.db")

	if err != nil {
		log.Fatal("Problem with opening the SQL DB Driver", err)
	}

	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`

	_, err = db.Exec(sqlStmt)

}
