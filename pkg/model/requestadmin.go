package model

import (
	"database/sql"

	"github.com/namay26/MVC-LMS/pkg/structs"
)

func RequestAdmin(db *sql.DB, user structs.User) bool {
	sqlquery := `UPDATE Users SET adminStatus='Pending' WHERE username=? AND userid=?`
	_, err := db.Exec(sqlquery, user.Username, user.Userid)
	if err != nil {
		return false
	}
	return true

}
