package controller

import (
	"log"
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/middleware"
	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/structs"
	"github.com/namay26/MVC-LMS/pkg/views"
	"golang.org/x/crypto/bcrypt"
)

func GetPage(w http.ResponseWriter, r *http.Request) {

	var data structs.Datasent
	views.Render(w, "index", data)
}

// Get login page

func LoginPage(w http.ResponseWriter, r *http.Request) {

	var message structs.PageMessage
	var err error

	message.Message, err = GetFlash(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	var data structs.Datasent
	data.Message = message
	views.Render(w, "login", data)
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

// Get Register Page

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
	var data structs.Datasent
	data.Message = message
	views.Render(w, "register", data)
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

		isAdmin, _, err := model.UserRegister(db, user)
		if err != nil {
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}

		userid, err2 := model.GetUserID(db, user.Username)
		if err2 != nil {
			log.Println(err2)
			http.Redirect(w, r, "/500", http.StatusSeeOther)
		}
		user.Userid = userid
		if isAdmin {
			user.IsAdmin = true
			middleware.SendCookie(w, user)
			http.Redirect(w, r, "/admin/home", http.StatusSeeOther)
		} else {
			middleware.SendCookie(w, user)
			http.Redirect(w, r, "/user/home", http.StatusSeeOther)
		}
	} else {
		SetFlash(w, r, "User already exists")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// Error Handles

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	var data structs.Datasent
	views.Render(w, "InternalServerError", data)
}

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	var data structs.Datasent
	views.Render(w, "PageNotFound", data)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "JWT",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
