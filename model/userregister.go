package model

import (
	"database/sql"
)

func UserRegister(db *sql.DB, username, password string) (bool, error) {
	_, err := db.Exec("INSERT INTO Users (username, pass) VALUES (?, ?)", username, password)
	if err != nil {
		return false, err
	}
	return true, nil
}
