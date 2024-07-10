package controller

import (
	"log"
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/model"
)

func AdminUpdate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	title := r.FormValue("title")
	author := r.FormValue("author")
	genre := r.FormValue("genre")
	db, _ := model.Connect()
	defer db.Close()

	updatesuccess, err := model.UpdateBook(db, id, title, author, genre)
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
