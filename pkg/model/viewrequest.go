package model

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/namay26/MVC-LMS/pkg/structs"
)

func ViewRequest(db *sql.DB) (structs.ListBookReq, error) {
	rows, err := db.Query("SELECT Users.userid, Users.username, books.id, books.title ,books.author, BookRequests.RequestDate FROM BookRequests JOIN books ON books.id=BookRequests.BookID JOIN Users ON BookRequests.UserID=Users.userid WHERE BookRequests.Status='Pending'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var listbookreq structs.ListBookReq
	for rows.Next() {
		var bh structs.BorrowHistory
		err := rows.Scan(&bh.UserID, &bh.Username, &bh.BookId, &bh.Title, &bh.Author, &bh.RequestDate)
		if err != nil {
			panic(err)
		}
		listbookreq.BorrowHistory = append(listbookreq.BorrowHistory, bh)
	}
	return listbookreq, nil

}

func AcceptRequest(db *sql.DB, userid string, bookid string) (bool, error) {
	uid, _ := strconv.Atoi(userid)
	bid, _ := strconv.Atoi(bookid)
	sqlquery := "UPDATE BookRequests SET Status = 'Approved', AcceptDate=NOW() WHERE BookID = ? AND UserID = ?"
	_, err := db.Exec(sqlquery, bid, uid)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
