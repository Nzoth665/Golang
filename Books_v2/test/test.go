package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	//"github.com/pressly/goose/v3/cmd/goose"
)

type product struct {
	id       int
	title    string
	year     int
	genre    int
	duration int
}

func main() {
	db, err := sql.Open("sqlite3", "films_db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from Films")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []product{}

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.title, &p.year, &p.genre, &p.duration)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Println(p.id, p.title, p.year, p.genre, p.duration)
	}
}
