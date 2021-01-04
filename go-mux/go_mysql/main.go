package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("my sql connection using go-lang")

	db, err := sql.Open("mysql", "root:lavanya123@tcp(127.0.0.1:3306)/golang")

	if err != nil {

		fmt.Println("Not connected")
		panic(err.Error())

	}

	defer db.Close()

	fmt.Println("Successfully connected to mysql DB")

	insert, err := db.Query("INSERT INTO posts VALUES(3,'Software engineer')")

	if err != nil {

		fmt.Println("Not inserted")
		panic(err.Error())

	}

	defer insert.Close()

	fmt.Println("Successfully inserted into emp table")

}
