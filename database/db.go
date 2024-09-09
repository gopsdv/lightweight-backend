package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

type Exercise struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Connect() error {
	envErr := godotenv.Load("/home/gdv/lightweight/.env")
    if envErr != nil {
        log.Fatalf("Error loading .env file")
    }

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBNAME"),
		AllowNativePasswords: true,
	}

	// Get a database handle.
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
	return nil
}


func CloseDB() {
    if DB != nil {
        DB.Close()
    }
}

