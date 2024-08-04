package model

import (
	"database/sql"

	"github.com/namay26/MVC-LMS/pkg/structs"
)

func UserRegister(db *sql.DB, person structs.User) (bool, bool, error) {

	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		return false, false, err
	}
	defer rows.Close()
	row2, err := db.Query("SELECT * FROM Users WHERE username = ?", person.Username)
	if err != nil {
		return false, false, err
	}
	defer row2.Close()
	if !row2.Next() {
		if !rows.Next() {
			_, err1 := db.Exec("INSERT INTO Users (username, pass, isAdmin, adminStatus) VALUES (?, ?, 1, 'isAdmin')", person.Username, person.Pass)
			if err1 != nil {
				return false, false, err1
			}
			return true, true, nil
		} else {
			_, err := db.Exec("INSERT INTO Users (username, pass) VALUES (?, ?)", person.Username, person.Pass)
			if err != nil {
				return false, false, err
			}
			return false, true, nil
		}
	} else {
		return false, false, nil
	}
}
