package model

import (
	"database/sql"

	"github.com/namay26/MVC-LMS/structs"
)

func RequestCheckout(db *sql.DB) (structs.ListBooks, error) {
	sqlquery := `SELECT * FROM books WHERE quantity > 0`
	rows, err := db.Query(sqlquery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var checkbook structs.ListBooks
	for rows.Next() {
		var book structs.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Quantity)
		if err != nil {
			panic(err)
		}
		checkbook.Books = append(checkbook.Books, book)
	}
	return checkbook, nil

}
