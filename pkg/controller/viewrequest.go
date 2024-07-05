package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetViewRequest(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	listusers, _ := model.ViewRequest(db)
	views.Render(w, "viewrequest", listusers)
}
