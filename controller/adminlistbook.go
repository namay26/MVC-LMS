package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/model"
	"github.com/namay26/MVC-LMS/views"
)

func AdminListBook(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	booklist, _ := model.GetBooks(db)
	views.Render(w, "adminlistbooks", booklist)
}
