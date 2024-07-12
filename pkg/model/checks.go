package model

import (
	"database/sql"
	"fmt"
	"strconv"
)

func CheckDuplicateBook(db *sql.DB, title string, author string) (bool, error) {
	rows, err := db.Query("SELECT id FROM books WHERE Title = ? AND Author = ?", title, author)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func CheckBorrowed(db *sql.DB, bookID string) (bool, error) {
	bid, _ := strconv.Atoi(bookID)
	rows, _ := db.Query("SELECT * FROM BookRequests WHERE BookID = ?", bid)
	if rows.Next() {
		return true, nil
	}
	return false, nil
}
