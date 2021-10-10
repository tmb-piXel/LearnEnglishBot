package main

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

func main() {
	runDB()
}

func runDB() {
	cfg := pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
	}
	conn, err := pgx.Connect(cfg)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT id_currency, charcode, name, numcode FROM currency;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	for rows.Next() {
		var id_currency int
		var name string
		var charcode string
		var numcode string
		err = rows.Scan(&id_currency, &charcode, &name, &numcode)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id_currency, charcode, name, numcode)
	}
}
