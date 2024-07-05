package controller

import (
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

	updatesuccess, _ := model.UpdateBook(db, id, title, author, genre)
	if updatesuccess {
		http.Redirect(w, r, "/admin/listbooks", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin/updatebook", http.StatusSeeOther)
	}
}
