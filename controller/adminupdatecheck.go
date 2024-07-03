package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/model"
	"github.com/namay26/MVC-LMS/views"
)

func AdminUpdateCheck(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")

	db, _ := model.Connect()
	defer db.Close()

	book, _ := model.GetBook(db, id)

	views.Render(w, "adminupdate", book)

}
