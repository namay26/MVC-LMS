package model

import (
	"database/sql"

	"github.com/namay26/MVC-LMS/structs"
)

func GetBook(db *sql.DB, id string) (structs.Book, error) {
	var book structs.Book
	err := db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
	if err != nil {
		return book, err
	}
	return book, nil
}
