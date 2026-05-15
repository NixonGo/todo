package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

var Db *sql.DB

const schema = `
CREATE TABLE scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT "",
	title VARCHAR(255) NOT NULL DEFAULT "",
	comment TEXT,
	repeat VARCHAR(128) NOT NULL DEFAULT ""
);

CREATE INDEX idx_date ON scheduler(date);
`

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)

	install := false

	if err != nil {
		install = true
	}

	Db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	if install {
		_, err = Db.Exec(schema)
		if err != nil {
			Db.Close()
			return err
		}
	}

	return nil
}
