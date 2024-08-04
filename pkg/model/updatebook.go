package model

import (
	"database/sql"
	"strconv"
)

func UpdateBook(db *sql.DB, id, title, author, genre, quant string) (bool, error) {
	i, _ := strconv.Atoi(id)
	q, _ := strconv.Atoi(quant)
	sqlquery := "UPDATE books SET title = ?, author = ?, genre = ?, quantity = ? WHERE id = ?"
	_, err := db.Exec(sqlquery, title, author, genre, q, i)

	if err != nil {
		return false, err
	}
	return true, nil
}
