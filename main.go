package main

import (
	"db_seminar/generator"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		"localhost", "5432", "postgres", "afimall", "12345", "disable"))

	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalln(err)
	}

	generator.CreateTables(db)
	generator.FillInDataBase(db)

	err = db.Close()

	if err != nil {
		fmt.Println(err)
	}
}
