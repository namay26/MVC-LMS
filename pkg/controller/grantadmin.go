package controller

import (
	"log"
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetGrantAdmin(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	requests, err := model.GrantAdmin(db)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	views.Render(w, "grantadmin", requests)
}

func GrantAdmin(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	r.ParseForm()
	userId := r.FormValue("userid")

	success, err := model.GrantAdminUpdate(db, userId)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	if !success {
		http.Redirect(w, r, "/admin/grantadmin", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/admin/grantadmin", http.StatusSeeOther)
}
