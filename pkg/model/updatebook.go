package model

import (
	"database/sql"
	"strconv"
)

func UpdateBook(db *sql.DB, id, title, author, genre string) (bool, error) {
	i, _ := strconv.Atoi(id)
	sqlquery := "UPDATE books SET title = ?, author = ?, genre = ? WHERE id = ?"
	_, err := db.Exec(sqlquery, title, author, genre, i)

	if err != nil {
		return false, err
	}
	return true, nil
}
