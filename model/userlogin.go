package model

import (
	"database/sql"
	"fmt"

	"github.com/namay26/MVC-LMS/structs"
)

func UserLogin(db *sql.DB, username string, password string) (bool, structs.User) {
	var user structs.User

	selectSql := `SELECT userid, username, pass, isAdmin FROM Users where username = ?`

	err := db.QueryRow(selectSql, username).Scan(&user.Userid, &user.Username, &user.Pass, &user.IsAdmin)
	if err != nil {
		return false, user
	}
	if password == user.Pass {
		fmt.Println("Password:", user.Pass)
		return true, user
	}
	return false, user
}
