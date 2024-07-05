package model

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/namay26/MVC-LMS/pkg/structs"
)

func GrantAdmin(db *sql.DB) (structs.ListUsers, error) {
	rows, err := db.Query("Select userid, username from Users where isAdmin=0 AND adminStatus='Pending'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var userlist structs.ListUsers
	for rows.Next() {
		var user structs.User
		err := rows.Scan(&user.Userid, &user.Username)
		if err != nil {
			log.Fatal(err)
		}
		userlist.Users = append(userlist.Users, user)
	}
	return userlist, nil
}

func GrantAdminUpdate(db *sql.DB, userId string) bool {

	uid, _ := strconv.Atoi(userId)
	_, err := db.Exec("Update Users set isAdmin=1, adminStatus='isAdmin' where userid=?", uid)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true

}
