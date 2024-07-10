package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/structs"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetAddBook(w http.ResponseWriter, r *http.Request) {
	var message structs.PageMessage
	var err error
	message.Message, err = GetFlash(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	views.Render(w, "addbook", message)
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
			http.Redirect(w, r, "/admin/listbooks", http.StatusSeeOther)
			SetFlash(w, r, "Book Added Successfully.")
		} else {
			fmt.Println("Error adding book")
		}
	}
}
