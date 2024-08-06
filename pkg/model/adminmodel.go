package model

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/namay26/MVC-LMS/pkg/types"
)

func AddBook(db *sql.DB, Title, Author, Genre, Quantity string) (bool, error) {
	quant, _ := strconv.Atoi(Quantity)
	_, err := db.Exec("INSERT INTO books(title, author, genre, quantity) VALUES(?, ?, ?, ?)", Title, Author, Genre, quant)
	if err != nil {
		return false, err
	}
	return true, nil
}

func AddQuantity(db *sql.DB, Title, Author, Quantity string) (bool, error) {
	quant, _ := strconv.Atoi(Quantity)
	_, err := db.Exec("UPDATE books SET quantity = quantity + ? WHERE title = ? AND author = ?", quant, Title, Author)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteBook(db *sql.DB, bookID string) (bool, error) {
	id, _ := strconv.Atoi(bookID)

	_, err := db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GrantAdmin(db *sql.DB) (types.ListUsers, error) {
	rows, err := db.Query("Select userid, username from Users where isAdmin=0 AND adminStatus='Pending'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var userlist types.ListUsers
	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.Userid, &user.Username)
		if err != nil {
			log.Fatal(err)
		}
		userlist.Users = append(userlist.Users, user)
	}
	return userlist, nil
}

func GrantAdminUpdate(db *sql.DB, userId string) (bool, error) {

	uid, _ := strconv.Atoi(userId)
	_, err := db.Exec("Update Users set isAdmin=1, adminStatus='isAdmin' where userid=?", uid)
	if err != nil {
		return false, err
	}
	return true, nil

}

func DenyAdminUpdate(db *sql.DB, userId string) (bool, error) {

	uid, _ := strconv.Atoi(userId)
	_, err := db.Exec("Update Users set isAdmin=0, adminStatus='NotRequested' where userid=?", uid)
	if err != nil {
		return false, err
	}
	return true, nil

}

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

func ViewRequest(db *sql.DB) (types.ListBookReq, error) {
	rows, err := db.Query("SELECT Users.userid, Users.username, books.id, books.title ,books.author, BookRequests.RequestDate FROM BookRequests JOIN books ON books.id=BookRequests.BookID JOIN Users ON BookRequests.UserID=Users.userid WHERE BookRequests.Status='Pending'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var listbookreq types.ListBookReq
	for rows.Next() {
		var bh types.BorrowHistory
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

func DenyRequest(db *sql.DB, userid string, bookid string) (bool, error) {
	uid, _ := strconv.Atoi(userid)
	bid, _ := strconv.Atoi(bookid)

	sqlquery := "UPDATE BookRequests SET Status = 'Returned' WHERE BookID = ? AND UserID = ?"
	_, err := db.Exec(sqlquery, bid, uid)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func Checkout(db *sql.DB, bookid string, userid float64) error {
	sqlquery := `INSERT INTO BookRequests (BookID, UserID, Status) VALUES (?, ?, "Pending")`
	_, err := db.Exec(sqlquery, bookid, userid)
	if err != nil {
		return err
	}
	return nil
}
