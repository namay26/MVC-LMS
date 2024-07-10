package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/views"
)

func AdminHome(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "adminhome", nil)
}
