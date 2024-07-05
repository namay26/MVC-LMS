package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetGrantAdmin(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	requests, _ := model.GrantAdmin(db)

	views.Render(w, "grantadmin", requests)
}

func GrantAdmin(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	r.ParseForm()
	userId := r.FormValue("userid")

	success := model.GrantAdminUpdate(db, userId)
	if !success {
		http.Redirect(w, r, "/admin/grantadmin", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/admin/grantadmin", http.StatusSeeOther)
}
