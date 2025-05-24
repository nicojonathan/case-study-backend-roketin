package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "workbench_user:workbenchpwd@tcp(localhost:3306)/db_short_movie_festival")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
