package model

import (
	"database/sql"

	"github.com/namay26/MVC-LMS/structs"
	"golang.org/x/crypto/bcrypt"
)

func UserFound(db *sql.DB, username string) bool {
	selectSql := `SELECT username FROM Users where username = ?`
	err := db.QueryRow(selectSql, username).Scan(&username)
	if err != nil {
		return false
	}
	return true
}

func PasswordMatch(db *sql.DB, username, password string) (bool, structs.User) {
	selectSql := `SELECT userid, username, pass, isAdmin FROM Users where username = ?`
	var user structs.User
	err := db.QueryRow(selectSql, username).Scan(&user.Userid, &user.Username, &user.Pass, &user.IsAdmin)

	if err != nil {
		return false, user
	}

	if err1 := bcrypt.CompareHashAndPassword(user.Pass, []byte(password)); err1 != nil {
		return false, user
	}

	return true, user
}
