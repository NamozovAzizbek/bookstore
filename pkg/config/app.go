package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "bookstore"
)

func Connect() *sql.DB {
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println(err)
	}

	// check db
	err = db.Ping()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return db
}
