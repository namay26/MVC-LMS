package controller

import (
	"log"
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/middleware"
	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetBorrowHistory(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()
	user := middleware.GetUser(r)
	borrowhistory, err := model.GetBorrowHistory(db, user)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	views.Render(w, "borrowhistory", borrowhistory)
}
