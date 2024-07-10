package controller

import (
	"log"
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

	checkBorrowed, err := model.CheckBorrowed(db, bookid)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	if checkBorrowed {
		log.Println("Book is borrowed!")
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
