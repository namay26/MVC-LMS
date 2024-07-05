package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/middleware"
	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetBorrowHistory(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()
	user := middleware.GetUser(r)
	borrowhistory, _ := model.GetBorrowHistory(db, user)
	views.Render(w, "borrowhistory", borrowhistory)
}
