package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func AdminUpdateCheck(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")

	db, _ := model.Connect()
	defer db.Close()

	Book, _ := model.GetBook(db, id)
	views.Render(w, "adminupdate", Book)

}
