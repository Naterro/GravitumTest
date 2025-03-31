package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var DB *sql.DB

func Connect(connStr string) {
	fmt.Println("Starting connection with Postgres Db")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Println("DB Ping Failed")
		time.Sleep(3 * time.Second)
		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("DB Connection started successfully")
	DB = db
}

func Close() {
	DB.Close()
}
