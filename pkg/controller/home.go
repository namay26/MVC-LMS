package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "home", nil)
}

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "PageNotFound", nil)
}
