package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func ListBooks(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	booklist, _ := model.GetBooks(db)
	views.Render(w, "listbooks", booklist)
}
