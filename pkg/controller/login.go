package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/middleware"
	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "JWT",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	views.Render(w, "login", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	username := r.FormValue("username")
	password := r.FormValue("password")

	userFound := model.UserFound(db, username)

	if userFound {
		passwordMatch, user := model.PasswordMatch(db, username, password)
		if passwordMatch {
			middleware.SendCookie(w, user)
			if user.IsAdmin {
				http.Redirect(w, r, "/admin/home", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/user/home", http.StatusSeeOther)
			}
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
