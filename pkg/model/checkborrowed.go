package model

import (
	"database/sql"
	"strconv"
)

func CheckBorrowed(db *sql.DB, bookID string) (bool, error) {
	bid, _ := strconv.Atoi(bookID)
	rows, _ := db.Query("SELECT * FROM BookRequests WHERE BookID = ?", bid)
	if rows.Next() {
		return true, nil
	}
	return false, nil
}
