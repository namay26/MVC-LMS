package controller

import (
	"log"
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/middleware"
	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/structs"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "JWT",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	var message structs.PageMessage
	var err error
	message.Message, err = GetFlash(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	views.Render(w, "login", message)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	username := r.FormValue("username")
	password := r.FormValue("password")

	userFound, err1 := model.UserFound(db, username)
	if err1 != nil {
		log.Println(err1)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	if userFound {
		passwordMatch, user, err := model.PasswordMatch(db, username, password)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/500", http.StatusSeeOther)
			return
		}
		if passwordMatch {
			middleware.SendCookie(w, user)
			if user.IsAdmin {
				http.Redirect(w, r, "/admin/home", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/user/home", http.StatusSeeOther)
			}
		} else {
			SetFlash(w, r, "Password doesn't match")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	} else {
		SetFlash(w, r, "User doesn't Exist.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
}
