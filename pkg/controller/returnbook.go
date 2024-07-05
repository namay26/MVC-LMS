package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/middleware"
	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetReturnBook(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	user := middleware.GetUser(r)

	returnbook, _ := model.GetReturnBook(db, user)
	views.Render(w, "returnbook", returnbook)

}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	bookid := r.FormValue("bookid")

	user := middleware.GetUser(r)

	model.ReturnBook(db, user, bookid)

	http.Redirect(w, r, "/user/returnbook", http.StatusSeeOther)
}
