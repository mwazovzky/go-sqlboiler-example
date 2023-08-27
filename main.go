package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-sqlboiler/models"
	"go-sqlboiler/services/config"
	"go-sqlboiler/services/database"
	"log"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TransactionWithClientAndCurrency struct {
	ID          uint        `boil:"transactions_id"`
	ClientID    uint        `boil:"clients_id"`
	ClientName  string      `boil:"clients_name"`
	CurrencyID  uint        `boil:"currencies_name"`
	CurrencyIso string      `boil:"currencies_iso"`
	Chain       null.String `boil:"transactions_chain"`
	Type        string      `boil:"transactions_type"`
	Status      string      `boil:"transactions_status"`
	Txid        string      `boil:"transactions_txid"`
	Vout        null.Uint   `boil:"transactions_vout"`
	Amount      uint64      `boil:"transactions_amount"`
	CreatedAt   null.Time   `boil:"transactions_created_at"`
	UpdatedAt   null.Time   `boil:"transactions_updated_at"`
}

func main() {
	cfg := config.Load()

	db := database.OpenDB(cfg)
	defer database.CloseDB(db)

	ctx := context.Background()

	// err := createClient(ctx, db)
	// handleErr(err)

	// txs, err := transactionsRawQuery(ctx, db)
	// handleErr(err)

	txs, err := transactionsQuery(ctx, db)
	handleErr(err)

	log.Printf("%#v\n", txs)
	log.Println(len(txs))
}

func transactionsQuery(ctx context.Context, db *sql.DB) ([]TransactionWithClientAndCurrency, error) {
	txs := []TransactionWithClientAndCurrency{}
	from := "2023-08-26 00:00:00"
	to := "2023-08-27 10:00:00"

	err := models.NewQuery(
		Select(
			"transactions.id as transactions_id",
			"clients.id as clients_id",
			"clients.name as clients_name",
			"currencies.id as currencies_id",
			"currencies.iso as currencies_iso",
			"transactions.chain as transactions_chain",
			"transactions.type as transactions_type",
			"transactions.status as transactions_status",
			"transactions.txid as transactions_txid",
			"transactions.vout as transactions_vout",
			"transactions.amount as transactions_amount",
			"transactions.created_at as transactions_created_at",
			"transactions.updated_at as transactions_updated_at",
		),
		From("transactions"),
		InnerJoin("clients on clients.id=transactions.client_id"),
		InnerJoin("currencies on currencies.id=transactions.currency_id"),
		Where("transactions.chain=?", "ethereum"),
		And("transactions.status=?", "cancelled"),
		And("transactions.amount > ?", 5000_0000),
		And("transactions.created_at >= ?", from),
		And("transactions.created_at <= ?", to),
	).Bind(ctx, db, &txs)

	if err != nil {
		return nil, err
	}

	return txs, nil
}

func createClient(ctx context.Context, db *sql.DB) error {
	count, err := models.Clients().Count(ctx, db)
	if err != nil {
		return err
	}

	name := fmt.Sprintf("Client #%d", count+1)

	c := &models.Client{Name: name}
	err = c.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}

	clients, err := models.Clients().All(ctx, db)
	if err != nil {
		return err
	}

	for _, c := range clients {
		log.Printf("%#v\n", *c)
	}

	return nil
}

func handleErr(err error) {
	if err != nil {
		log.Fatal("error:", err)
	}
}
