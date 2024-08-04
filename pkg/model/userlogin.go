package model

import (
	"database/sql"
	"log"

	"github.com/namay26/MVC-LMS/pkg/structs"
	"golang.org/x/crypto/bcrypt"
)

func UserFound(db *sql.DB, usernm string) (bool, error) {
	selectSql := `SELECT username FROM Users where username = ?`
	rows, err := db.Query(selectSql, usernm)
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}

func PasswordMatch(db *sql.DB, username, password string) (bool, structs.User, error) {
	selectSql := `SELECT userid, username, pass, isAdmin FROM Users where username = ?`
	var user structs.User
	err := db.QueryRow(selectSql, username).Scan(&user.Userid, &user.Username, &user.Pass, &user.IsAdmin)

	if err != nil {
		return false, user, err
	}

	if err1 := bcrypt.CompareHashAndPassword(user.Pass, []byte(password)); err1 != nil {
		return false, user, nil
	}

	return true, user, nil
}

func GetUserID(db *sql.DB, username string) (float64, error) {
	selectSql := `SELECT userid FROM Users where username = ?`
	var userid float64
	err := db.QueryRow(selectSql, username).Scan(&userid)
	if err != nil {
		return 0, err
	}
	return userid, nil
}
