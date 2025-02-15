package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB

func InitDB() *sql.DB {
	var err error
	server := os.Getenv("DB_SERVER")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	connString := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s;encrypt=disable",
		server, port, user, password, database)

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error verifying connection: %v", err)
	}

	fmt.Println("Connected to MSSQL!")
	return db
}
