package main

import (
	"db_seminar/newGenerator"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		"localhost", "5432", "postgres", "oait", "12345", "disable"))

	if err != nil {
		fmt.Println(err)
		fmt.Println("хуйхуй")
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	err = db.Ping()

	if err != nil {
		log.Fatalln(err)
	}

	//newGenerator.CreateTables(db)

	//newGenerator.FillInDataBase(db)

	newGenerator.GenerateData(db)

	fmt.Println("all tables are created and filled")
}
