package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/structs"
	"github.com/namay26/MVC-LMS/pkg/views"
)

// Admin Dashboard
func AdminHome(w http.ResponseWriter, r *http.Request) {
	var data structs.Datasent
	views.Render(w, "adminhome", data)
}

// Add Book Functions
func GetAddBook(w http.ResponseWriter, r *http.Request) {
	var message structs.PageMessage
	var err error
	message.Message, err = GetFlash(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var data structs.Datasent
	data.Message = message
	views.Render(w, "addbook", data)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Title := r.FormValue("title")
	Author := r.FormValue("author")
	Genre := r.FormValue("genre")
	Quantity := r.FormValue("quantity")

	db, _ := model.Connect()
	defer db.Close()

	CheckDuplicates, err := model.CheckDuplicateBook(db, Title, Author)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	if CheckDuplicates {
		AddQuantity, err2 := model.AddQuantity(db, Title, Author, Quantity)
		if err2 != nil {
			log.Println(err2)
			http.Redirect(w, r, "/500", http.StatusSeeOther)
			return
		}
		if AddQuantity {
			SetFlash(w, r, "Book Quantity updated.")
			http.Redirect(w, r, "/admin/listbooks", http.StatusSeeOther)
		} else {
			fmt.Println("Error adding book")
		}
	} else {
		AddSuccessful, err3 := model.AddBook(db, Title, Author, Genre, Quantity)
		if err3 != nil {
			log.Println(err3)
			http.Redirect(w, r, "/500", http.StatusSeeOther)
			return
		}
		if AddSuccessful {
			SetFlash(w, r, "Book Added Successfully.")
			http.Redirect(w, r, "/admin/listbooks", http.StatusSeeOther)
		} else {
			fmt.Println("Error adding book")
		}
	}
}

// Admin List Book Page

func GetAdminListBook(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	booklist, err2 := model.GetBooks(db)
	if err2 != nil {
		log.Println(err2)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var message structs.PageMessage
	var err error
	message.Message, err = GetFlash(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var data structs.Datasent
	data.Message = message
	data.Results = booklist
	views.Render(w, "adminlistbooks", data)
}

// Admin Book Update

func GetAdminUpdate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")

	db, _ := model.Connect()
	defer db.Close()

	Book, err2 := model.GetBook(db, id)
	if err2 != nil {
		log.Println(err2)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var message structs.PageMessage
	var err error
	message.Message, err = GetFlash(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var data structs.Datasent
	data.Results = Book
	data.Message = message
	views.Render(w, "adminupdate", data)

}

func AdminUpdate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	title := r.FormValue("title")
	author := r.FormValue("author")
	genre := r.FormValue("genre")
	quant := r.FormValue("quantity")

	db, _ := model.Connect()
	defer db.Close()

	updatesuccess, err := model.UpdateBook(db, id, title, author, genre, quant)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	if updatesuccess {
		SetFlash(w, r, "Book Updated Successfully!")
		http.Redirect(w, r, "/admin/listbooks", http.StatusSeeOther)
	} else {
		SetFlash(w, r, "Book couldn't be updated. Please Try again!")
		http.Redirect(w, r, "/admin/updatebook", http.StatusSeeOther)
	}
}

// Admin Delete Book

func GetDeleteBook(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	booklist, err2 := model.GetBooks(db)
	if err2 != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		panic(err2)
	}
	var message structs.PageMessage
	var err error
	message.Message, err = GetFlash(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var data structs.Datasent
	data.Message = message
	data.Results = booklist
	views.Render(w, "deletebook", data)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bookid := r.FormValue("id")

	db, _ := model.Connect()
	defer db.Close()

	checkBorrowed, err := model.CheckBorrowed(db, bookid)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	if checkBorrowed {
		SetFlash(w, r, "Book is borrowed, cannot delete!")
		http.Redirect(w, r, "/admin/deletebook", http.StatusSeeOther)
		return
	}
	deletesuccess, err1 := model.DeleteBook(db, bookid)
	if err1 != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	if deletesuccess {
		SetFlash(w, r, "Book deleted Successfully!")
		http.Redirect(w, r, "/admin/deletebook", http.StatusSeeOther)
	} else {
		SetFlash(w, r, "Book couldnt be deleted")
		http.Redirect(w, r, "/admin/deletebook", http.StatusSeeOther)
	}

}

// Admin Grant Admin

func GetGrantAdmin(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	requests, err2 := model.GrantAdmin(db)
	if err2 != nil {
		log.Println(err2)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var message structs.PageMessage
	var err error
	message.Message, err = GetFlash(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var data structs.Datasent
	data.Results = requests
	data.Message = message
	views.Render(w, "grantadmin", data)
}

func GrantAdmin(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	r.ParseForm()
	userId := r.FormValue("userid")

	success, err := model.GrantAdminUpdate(db, userId)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	if !success {
		SetFlash(w, r, "Admin Grant Failed!")
		http.Redirect(w, r, "/admin/grantadmin", http.StatusSeeOther)
		return
	}
	SetFlash(w, r, "Admin Granted Successfully!")
	http.Redirect(w, r, "/admin/grantadmin", http.StatusSeeOther)
}

// Admin Book Requests

func GetViewRequest(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	listusers, _ := model.ViewRequest(db)
	var message structs.PageMessage
	var err error
	message.Message, err = GetFlash(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var data structs.Datasent
	data.Message = message
	data.Results = listusers
	views.Render(w, "viewrequest", data)
}

func ViewRequest(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	r.ParseForm()
	userid := r.FormValue("userid")
	bookid := r.FormValue("bookid")
	acptreq, err := model.AcceptRequest(db, userid, bookid)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	if acptreq {
		SetFlash(w, r, "Request Accepted")
		http.Redirect(w, r, "/admin/viewrequest", http.StatusSeeOther)
	} else {
		fmt.Println("Error")
	}

}
