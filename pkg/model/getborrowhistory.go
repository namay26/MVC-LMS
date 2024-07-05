package model

import (
	"database/sql"

	"github.com/namay26/MVC-LMS/pkg/structs"
)

func GetBorrowHistory(db *sql.DB, user structs.User) (structs.ListBookReq, error) {
	rows, err := db.Query("SELECT BookRequests.RequestID,BookRequests.Status, books.title, books.author, books.genre, BookRequests.RequestDate FROM BookRequests JOIN books ON books.id=BookRequests.BookID WHERE BookRequests.UserID=?", user.Userid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var listbookreq structs.ListBookReq
	for rows.Next() {
		var bh structs.BorrowHistory
		err := rows.Scan(&bh.BookId, &bh.Status, &bh.Title, &bh.Author, &bh.Genre, &bh.RequestDate)
		if err != nil {
			panic(err)
		}
		listbookreq.BorrowHistory = append(listbookreq.BorrowHistory, bh)
	}
	return listbookreq, nil
}
