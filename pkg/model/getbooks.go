package model

import (
	"database/sql"

	"github.com/namay26/MVC-LMS/pkg/structs"
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

func GetBook(db *sql.DB, id string) (structs.Book, error) {
	var Book structs.Book
	err := db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&Book.ID, &Book.Title, &Book.Author, &Book.Genre, &Book.Quantity)
	if err != nil {
		return Book, err
	}
	return Book, nil
}
