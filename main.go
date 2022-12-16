package main

import (
	"db_seminar/generator"
	"encoding/base64"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		"localhost", "5432", "postgres", "testest", "12345", "disable"))

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

	generator.CreateTables(db)

	generator.FillInDataBase(db)

	generator.GenerateData(db)

	s := "i am string"
	se := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(se)
	fmt.Println("all tables are created and filled")
}
