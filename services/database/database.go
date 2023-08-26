package database

import (
	"database/sql"
	"go-sqlboiler/services/config"
	"log"
)

func OpenDB(cfg config.Config) *sql.DB {
	dsn := cfg.FormatDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open DB connection:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB:", err)
	}

	log.Printf("Opened DB connection: %s/%s", cfg.Host, cfg.Database)

	return db
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal("Failed to close DB connection:", err)
	}

	log.Println("Closed DB connection")
}
