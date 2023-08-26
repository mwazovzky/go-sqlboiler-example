package main

import (
	"go-sqlboiler/services/config"
	"go-sqlboiler/services/database"
)

func main() {
	cfg := config.Load()

	db := database.OpenDB(cfg)
	defer database.CloseDB(db)
}
