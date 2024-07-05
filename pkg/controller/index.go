package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetPage(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "JWT",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	views.Render(w, "index", nil)
}
