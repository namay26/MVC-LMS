package controller

import (
	"log"
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/structs"
	"github.com/namay26/MVC-LMS/pkg/views"
	"golang.org/x/crypto/bcrypt"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
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

	views.Render(w, "register", message)
}

func Register(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	userFound, err1 := model.UserFound(db, r.FormValue("username"))
	if err1 != nil {
		log.Println(err1)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	if !userFound {
		password, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)

		user := structs.User{
			Username: r.FormValue("username"),
			Pass:     password,
		}

		_, err := model.UserRegister(db, user)
		if err != nil {

			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		SetFlash(w, r, "User already exists")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
