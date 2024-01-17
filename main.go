package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Transaction struct {
	ID     string
	Amount float64
	Status string
}

func CreateTransaction(db *sql.DB, t *Transaction) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO transactions(id, amount, status) VALUES($1, $2, 'created')", t.ID, t.Amount)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func ReverseTransaction(db *sql.DB, t *Transaction) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE transactions SET status='reversed' WHERE id=$1", t.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func main() {
	connStr := "user=postgres password=postgres dbname=reversed sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() 

	t := &Transaction{ID: "3010712345678901", Amount: 100.0}

	err = CreateTransaction(db, t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction created.")

	err = ReverseTransaction(db, t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction reversed.")
}

