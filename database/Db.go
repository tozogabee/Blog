package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "user.db"
)

func Connect() *sql.DB {
	connInfo := fmt.Sprintf("host=%s port=%d "+
		"dbname=%s sslmode=disable",
		host, port, dbname)

	db, err := sql.Open("sqlite3", connInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to db!")
	return db
}

func CloseConnection(db *sql.DB) {
	defer db.Close()
}
