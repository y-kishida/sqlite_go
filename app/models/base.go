package models

import (
	"database/sql"
	"fmt"
	"log"
	"udemy-golang/go_crud_app/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

const tableNameUser = "users"

func init() {
	Db, err := sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)
}
