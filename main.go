package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// postgresql://yusufhalabi@localhost:5432/quran

func main() {
	// start by adding the connection to the db locally
	// pgx.Connect returns a pointer to a pgx.Conn object
	conn, err := pgx.Connect(context.Background(), "postgresql://yusufhalabi@localhost:5432/quran") // Connect to the database

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background()) // Close the connection when done

	// now run a query
	var surah_name string
	var ayah_num int
	var ayah string

	err = conn.QueryRow(context.Background(), "SElECT surah_name_en, surah_no, ayah_ar FROM quran LIMIT 1;").Scan(&surah_name, &ayah_num, &ayah) // pass in reference to res to load in result of query
	if err != nil {
		fmt.Fprintf(os.Stderr, "query failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(surah_name, ayah_num, ayah)

}
