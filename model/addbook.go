package model

import (
	"database/sql"
	"strconv"
)

func AddBook(db *sql.DB, Title, Author, Genre, Quantity string) (bool, error) {
	quant, _ := strconv.Atoi(Quantity)
	_, err := db.Exec("INSERT INTO books(title, author, genre, quantity) VALUES(?, ?, ?, ?)", Title, Author, Genre, quant)
	if err != nil {
		return false, err
	}
	return true, nil
}
