package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetDeleteBook(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	booklist, err := model.GetBooks(db)
	if err != nil {
		panic(err)
	}
	views.Render(w, "deletebook", booklist)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bookid := r.FormValue("id")

	db, _ := model.Connect()
	defer db.Close()

	checkBorrowed, _ := model.CheckBorrowed(db, bookid)
	if checkBorrowed {
		http.Redirect(w, r, "/admin/deletebook", http.StatusSeeOther)
		return
	}
	deletesuccess, err := model.DeleteBook(db, bookid)
	if err != nil {
		panic(err)
	}
	if deletesuccess {
		http.Redirect(w, r, "/admin/deletebook", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin/deletebook", http.StatusSeeOther)
	}

}
