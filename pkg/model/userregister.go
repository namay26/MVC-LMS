package model

import (
	"database/sql"

	"github.com/namay26/MVC-LMS/pkg/structs"
)

func UserRegister(db *sql.DB, person structs.User) (bool, error) {
	_, err := db.Exec("INSERT INTO Users (username, pass) VALUES (?, ?)", person.Username, person.Pass)
	if err != nil {
		return false, err
	}
	return true, nil
}
