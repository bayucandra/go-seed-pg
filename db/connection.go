package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
	"os"
)

var DBConn *sql.DB

func Init() {

	var err error
	DBConn, err = connect()

	if err != nil {
		log.Fatalf("Error connecting Database: %v", err)
	}
}

func connect() (db *sql.DB, err error) {
	var (
		host     = os.Getenv("PG_HOST")
		port     = os.Getenv("PG_PORT")
		user     = os.Getenv("PG_USER")
		password = os.Getenv("PG_PASSWORD")
		dbname   = os.Getenv("PG_DBNAME")
		sslmode  = os.Getenv("PG_SSLMODE")
	)

	if host == "docker-host" && os.Getenv("GO_ENV") == "development" {
		host = os.Getenv("DOCKER_HOST")
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode,
	)

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		return
	}

	err = db.Ping()

	if err != nil {
		return
	}
	return
}