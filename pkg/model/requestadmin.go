package model

import (
	"database/sql"

	"github.com/namay26/MVC-LMS/pkg/structs"
)

func RequestAdmin(db *sql.DB, user structs.User) (bool, error) {

	admincheck := `SELECT adminStatus FROM Users WHERE username=? AND userid=?`
	row := db.QueryRow(admincheck, user.Username, user.Userid)
	var adminStatus string
	err2 := row.Scan(&adminStatus)
	if err2 != nil {
		return false, err2
	}

	if adminStatus == "Pending" {
		return false, nil
	}

	sqlquery := `UPDATE Users SET adminStatus='Pending' WHERE username=? AND userid=?`
	_, err := db.Exec(sqlquery, user.Username, user.Userid)
	if err != nil {
		return false, err
	}
	return true, nil

}

func CheckRequest(db *sql.DB, user structs.User) (bool, error) {

	admincheck := `SELECT adminStatus FROM Users WHERE username=? AND userid=?`
	row := db.QueryRow(admincheck, user.Username, user.Userid)
	var adminStatus string
	err2 := row.Scan(&adminStatus)
	if err2 != nil {
		return false, err2
	}
	if adminStatus == "Pending" {
		return true, nil
	}
	return false, nil
}
