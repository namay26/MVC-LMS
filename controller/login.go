package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/model"
	"github.com/namay26/MVC-LMS/views"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "login", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	db, _ := model.Connect()
	defer db.Close()

	loginSuccess, _ := model.UserLogin(db, username, password)

	if loginSuccess {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		views.Render(w, "/register", "Invalid Username or Password")
	}
}
