package model

import "database/sql"

func Checkout(db *sql.DB, bookid string, userid string) error {
	sqlquery := `INSERT INTO BookRequests (bookid, userid) VALUES (?, ?)`
	_, err := db.Exec(sqlquery, bookid, userid)
	if err != nil {
		return err
	}
	return nil
}
