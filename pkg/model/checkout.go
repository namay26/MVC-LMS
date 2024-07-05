package model

import (
	"database/sql"
)

func Checkout(db *sql.DB, bookid string, userid float64) error {
	sqlquery := `INSERT INTO BookRequests (BookID, UserID, Status) VALUES (?, ?, "Pending")`
	_, err := db.Exec(sqlquery, bookid, userid)
	if err != nil {
		return err
	}
	return nil
}
