package model

import (
	"database/sql"

	"github.com/namay26/MVC-LMS/structs"
)

func GetBooks(db *sql.DB) (structs.ListBooks, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var listbook structs.ListBooks
	for rows.Next() {
		var book structs.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Quantity)
		if err != nil {
			panic(err)
		}
		listbook.Books = append(listbook.Books, book)
	}
	return listbook, nil
}
