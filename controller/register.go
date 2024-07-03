package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/model"
	"github.com/namay26/MVC-LMS/views"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "register", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	db, _ := model.Connect()
	defer db.Close()
	//Hashing password
	registerSuccess, _ := model.UserRegister(db, username, password)

	if registerSuccess {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
}
