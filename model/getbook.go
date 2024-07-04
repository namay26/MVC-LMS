package model

import (
	"database/sql"

	"github.com/namay26/MVC-LMS/structs"
)

func GetBook(db *sql.DB, id string) (structs.Book, error) {
	var Book structs.Book
	err := db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&Book.ID, &Book.Title, &Book.Author, &Book.Genre, &Book.Quantity)
	if err != nil {
		return Book, err
	}
	return Book, nil
}
