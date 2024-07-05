package controller

import (
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
	views.Render(w, "register", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	password, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
	user := structs.User{
		Username: r.FormValue("username"),
		Pass:     password,
	}

	_, err := model.UserRegister(db, user)
	if err != nil {
		views.Render(w, "register", "Error registering user")
		return
	}
	views.Render(w, "login", "User registered successfully")
}
