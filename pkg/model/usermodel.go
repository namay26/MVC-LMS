package model

import (
	"database/sql"
	"strconv"

	"github.com/namay26/MVC-LMS/pkg/types"
)

func GetBorrowHistory(db *sql.DB, user types.User) (types.ListBookReq, error) {
	rows, err := db.Query("SELECT BookRequests.RequestID,BookRequests.Status, books.title, books.author, books.genre, BookRequests.RequestDate FROM BookRequests JOIN books ON books.id=BookRequests.BookID WHERE BookRequests.UserID=?", user.Userid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var listbookreq types.ListBookReq
	for rows.Next() {
		var bh types.BorrowHistory
		err := rows.Scan(&bh.BookId, &bh.Status, &bh.Title, &bh.Author, &bh.Genre, &bh.RequestDate)
		if err != nil {
			panic(err)
		}
		listbookreq.BorrowHistory = append(listbookreq.BorrowHistory, bh)
	}
	return listbookreq, nil
}

func RequestAdmin(db *sql.DB, user types.User) (bool, error) {

	admincheck := `SELECT adminStatus FROM Users WHERE username=? AND userid=?`
	row := db.QueryRow(admincheck, user.Username, user.Userid)
	var adminStatus string
	err2 := row.Scan(&adminStatus)
	if err2 != nil {
		return false, err2
	}

	if adminStatus == "Pending" {
		return false, nil
	}

	sqlquery := `UPDATE Users SET adminStatus='Pending' WHERE username=? AND userid=?`
	_, err := db.Exec(sqlquery, user.Username, user.Userid)
	if err != nil {
		return false, err
	}
	return true, nil

}

func CheckRequest(db *sql.DB, user types.User) (bool, error) {

	admincheck := `SELECT adminStatus FROM Users WHERE username=? AND userid=?`
	row := db.QueryRow(admincheck, user.Username, user.Userid)
	var adminStatus string
	err2 := row.Scan(&adminStatus)
	if err2 != nil {
		return false, err2
	}
	if adminStatus == "Pending" {
		return true, nil
	}
	return false, nil
}

func GetReturnBook(db *sql.DB, user types.User) (types.ListBookReq, error) {
	rows, err := db.Query("SELECT books.id, books.title, books.author, books.genre, BookRequests.RequestDate, BookRequests.AcceptDate FROM BookRequests JOIN books ON books.id=BookRequests.BookID WHERE BookRequests.UserID=? AND BookRequests.Status='Approved'", user.Userid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var returnbook types.ListBookReq
	for rows.Next() {
		var bh types.BorrowHistory
		err := rows.Scan(&bh.BookId, &bh.Author, &bh.Title, &bh.Genre, &bh.RequestDate, &bh.AcceptDate)
		if err != nil {
			panic(err)
		}
		returnbook.BorrowHistory = append(returnbook.BorrowHistory, bh)
	}
	return returnbook, nil
}

func ReturnBook(db *sql.DB, user types.User, bookid string) error {
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

func RequestCheckout(db *sql.DB, user types.User) (types.ListBooks, error) {

	sqlquery := `SELECT * FROM books WHERE id NOT IN (SELECT BookID FROM BookRequests WHERE (UserID = ?) AND (Status='Pending' OR Status='Approved')) AND quantity > 0`

	rows, err := db.Query(sqlquery, user.Userid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var checkbook types.ListBooks
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Quantity)
		if err != nil {
			panic(err)
		}
		checkbook.Books = append(checkbook.Books, book)
	}
	return checkbook, nil

}
