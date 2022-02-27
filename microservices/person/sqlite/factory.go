package sqlite

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() *sql.DB {
	os.Remove("./rtdd.db")

	db, err := sql.Open("sqlite3", "./rtdd.db")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func ConfigConnection(db *sql.DB) {
	tablePerson := `
		CREATE TABLE person (
			id integer not null primary key autoincrement,
			name text,
			age integer
		);


		INSERT INTO person(name, age) VALUES ('Vinicius', 24);
		INSERT INTO person(name, age) VALUES ('Kevin', 23);
	`

	_, err := db.Exec(tablePerson)

	if err != nil {
		log.Fatal(err)
	}
}
