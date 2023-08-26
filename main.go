package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-sqlboiler/models"
	"go-sqlboiler/services/config"
	"go-sqlboiler/services/database"
	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	cfg := config.Load()

	db := database.OpenDB(cfg)
	defer database.CloseDB(db)

	ctx := context.Background()

	count, err := models.Clients().Count(ctx, db)
	handleErr(err)
	fmt.Println(count)

	name := fmt.Sprintf("Client #%d", count+1)

	c := &models.Client{Name: name}
	err = c.Insert(ctx, db, boil.Infer())
	handleErr(err)

	clients, err := models.Clients().All(ctx, db)
	handleErr(err)
	for _, c := range clients {
		fmt.Printf("%#v\n", *c)
		fmt.Println()
	}
}

func insertClient(db *sql.DB, c *models.Client) error {
	return c.Insert(context.Background(), db, boil.Infer())
}

func handleErr(err error) {
	if err != nil {
		log.Fatal("error:", err)
	}
}
