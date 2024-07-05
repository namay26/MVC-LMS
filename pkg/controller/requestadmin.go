package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/middleware"
	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetRequestAdmin(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "requestadmin", nil)
}

func RequestAdmin(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	user := middleware.GetUser(r)

	requestadmin := model.RequestAdmin(db, user)
	if !requestadmin {
		views.Render(w, "requestadmin", "Request failed")
		return
	} else {
		views.Render(w, "requestadmin", "Request sent")
		return
	}
}
