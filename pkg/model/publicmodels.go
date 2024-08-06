package model

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/namay26/MVC-LMS/pkg/types"
	"golang.org/x/crypto/bcrypt"
)

func CheckDuplicateBook(db *sql.DB, title string, author string) (bool, error) {
	rows, err := db.Query("SELECT id FROM books WHERE Title = ? AND Author = ?", title, author)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func CheckBorrowed(db *sql.DB, bookID string) (bool, error) {
	bid, _ := strconv.Atoi(bookID)
	rows, _ := db.Query("SELECT * FROM BookRequests WHERE BookID = ?", bid)
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func GetBooks(db *sql.DB) (types.ListBooks, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var listbook types.ListBooks
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Quantity)
		if err != nil {
			panic(err)
		}
		listbook.Books = append(listbook.Books, book)
	}
	return listbook, nil
}

func GetBook(db *sql.DB, id string) (types.Book, error) {
	var Book types.Book
	err := db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&Book.ID, &Book.Title, &Book.Author, &Book.Genre, &Book.Quantity)
	if err != nil {
		return Book, err
	}
	return Book, nil
}

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

func PasswordMatch(db *sql.DB, username, password string) (bool, types.User, error) {
	selectSql := `SELECT userid, username, pass, isAdmin FROM Users where username = ?`
	var user types.User
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

func UserRegister(db *sql.DB, person types.User) (bool, bool, error) {

	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		return false, false, err
	}
	defer rows.Close()
	row2, err := db.Query("SELECT * FROM Users WHERE username = ?", person.Username)
	if err != nil {
		return false, false, err
	}
	defer row2.Close()
	if !row2.Next() {
		if !rows.Next() {
			_, err1 := db.Exec("INSERT INTO Users (username, pass, isAdmin, adminStatus) VALUES (?, ?, 1, 'isAdmin')", person.Username, person.Pass)
			if err1 != nil {
				return false, false, err1
			}
			return true, true, nil
		} else {
			_, err := db.Exec("INSERT INTO Users (username, pass) VALUES (?, ?)", person.Username, person.Pass)
			if err != nil {
				return false, false, err
			}
			return false, true, nil
		}
	} else {
		return false, false, nil
	}
}
