package main

import (
	"context"
	"database/sql"
	"go-sqlboiler/models"
	"go-sqlboiler/services/config"
	"go-sqlboiler/services/database"
	"log"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const count = 100

type FakeTransaction struct {
	Txid   string `faker:"uuid_hyphenated"`
	Vout   uint   `faker:"boundary_start=0, boundary_end=10"`
	Amount uint64 `faker:"boundary_start=1, boundary_end=10000"`
}

func main() {
	cfg := config.Load()

	db := database.OpenDB(cfg)
	defer database.CloseDB(db)

	ctx := context.Background()

	clients, err := models.Clients().All(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	currencies, err := models.Currencies().All(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	deleteTransactions(ctx, db)

	// create tx only for crypto currencies
	for _, client := range clients {
		for _, currency := range currencies {
			if currency.Chain.Valid {
				seedTransactions(ctx, db, client, currency)
			}
		}
	}
}

func seedTransactions(ctx context.Context, db *sql.DB, client *models.Client, currency *models.Currency) error {
	for i := 0; i < count; i++ {
		err := createTransaction(ctx, db, client.ID, currency.ID, currency.Chain.String)
		if err != nil {
			log.Println("createTransaction error", err)
		}
	}

	return nil
}

func createTransaction(ctx context.Context, db *sql.DB, client uint, currency uint, chain string) error {
	types := []string{"deposit", "withdrawal"}
	statuses := []string{"created", "confirmed", "cancelled"}
	tm := getRandomTime()

	ftx := FakeTransaction{}
	err := faker.FakeData(&ftx)
	if err != nil {
		return err
	}

	tx := &models.Transaction{
		ClientID:   client,
		CurrencyID: currency,
		Chain:      null.StringFrom(chain),
		Type:       getRandomArrayItem(types),
		Status:     getRandomArrayItem(statuses),
		Txid:       ftx.Txid,
		Vout:       null.UintFrom(ftx.Vout),
		Amount:     ftx.Amount * 10000,
		CreatedAt:  null.TimeFrom(tm),
	}

	return tx.Insert(ctx, db, boil.Infer())
}

func deleteTransactions(ctx context.Context, db *sql.DB) error {
	_, err := models.Transactions().DeleteAll(ctx, db)
	return err
}

func getRandomArrayItem(items []string) string {
	index := rand.Intn(len(items))
	return items[index]
}

func getRandomTime() time.Time {
	interval := time.Duration(-14400 * time.Minute)
	from := time.Now().Add(interval).Unix()
	to := time.Now().Unix()
	d := rand.Int63n(to - from)

	return time.Unix(from+d, 0)
}
