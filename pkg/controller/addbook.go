package controller

import (
	"fmt"
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetAddBook(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "addbook", nil)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Title := r.FormValue("title")
	Author := r.FormValue("author")
	Genre := r.FormValue("genre")
	Quantity := r.FormValue("quantity")

	db, _ := model.Connect()
	defer db.Close()

	CheckDuplicates, _ := model.CheckDuplicateBook(db, Title, Author)

	if CheckDuplicates {
		AddQuantity, _ := model.AddQuantity(db, Title, Author, Quantity)
		if AddQuantity {
			http.Redirect(w, r, "/admin/listbooks", http.StatusSeeOther)
		} else {
			fmt.Println("Error adding book")
		}
	} else {
		AddSuccessful, _ := model.AddBook(db, Title, Author, Genre, Quantity)
		if AddSuccessful {
			http.Redirect(w, r, "/admin/listbooks", http.StatusSeeOther)
		} else {
			fmt.Println("Error adding book")
		}
	}
}
