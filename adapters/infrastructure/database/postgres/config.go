package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var instance *sql.DB

func init() {
	var (
		host    = os.Getenv("DATABASE_HOST")
		port    = os.Getenv("DATABASE_PORT")
		user    = os.Getenv("DATABASE_USER")
		pass    = os.Getenv("DATABASE_PASS")
		dbname  = os.Getenv("DATABASE_NAME")
		connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, pass, dbname)
	)

	if os.Getenv("ENV") == "local" {
		connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error while opening '%s' postgresql database: %s\n", dbname, err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error while ping '%s' postgresql database: %s\n", dbname, err.Error())
	}

	executeMigrations(db)

	instance = db
}

func GetInstance() *sql.DB {
	return instance
}
