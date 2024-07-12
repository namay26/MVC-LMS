package model

import (
	"database/sql"
	"strconv"

	"github.com/namay26/MVC-LMS/pkg/structs"
)

func GetReturnBook(db *sql.DB, user structs.User) (structs.ListBookReq, error) {
	rows, err := db.Query("SELECT books.id, books.title, books.author, books.genre, BookRequests.RequestDate, BookRequests.AcceptDate FROM BookRequests JOIN books ON books.id=BookRequests.BookID WHERE BookRequests.UserID=? AND BookRequests.Status='Approved'", user.Userid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var returnbook structs.ListBookReq
	for rows.Next() {
		var bh structs.BorrowHistory
		err := rows.Scan(&bh.BookId, &bh.Author, &bh.Title, &bh.Genre, &bh.RequestDate, &bh.AcceptDate)
		if err != nil {
			panic(err)
		}
		returnbook.BorrowHistory = append(returnbook.BorrowHistory, bh)
	}
	return returnbook, nil
}

func ReturnBook(db *sql.DB, user structs.User, bookid string) error {
	bid, err := strconv.Atoi(bookid)
	if err != nil {
		return err
	}
	sqlquery := "UPDATE BookRequests SET Status='Returned', ReturnDate=NOW() WHERE BookID=? AND UserID=?"
	_, err1 := db.Exec(sqlquery, bid, user.Userid)
	if err1 != nil {
		return err1
	}
	return err1
}
