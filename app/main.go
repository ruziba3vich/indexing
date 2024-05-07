package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	loader "github.com/ruziba3vich/dummy_people"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Dost0n1k", "indexing")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbName := "people"

	name := "../internal/db/" + dbName + ".sql"
	sqlFile, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(sqlFile))
	log.Println(string(sqlFile))
	if err != nil {
		log.Fatal(err)
	}

	ok, err := loader.LoadFromDatabase(db)
	if err != nil {
		log.Println(err)
	}
	if ok {
		fmt.Println("ok")
	}
}
