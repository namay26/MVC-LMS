package controller

import (
	"log"
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

	err := model.ReturnBook(db, user, bookid)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/user/returnbook", http.StatusSeeOther)
}
