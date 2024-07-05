package model

import (
	"database/sql"
	"strconv"
)

func DeleteBook(db *sql.DB, bookID string) (bool, error) {
	id, _ := strconv.Atoi(bookID)
	_, err := db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		return false, err
	}
	return true, nil
}
