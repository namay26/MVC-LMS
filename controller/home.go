package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "home", nil)
}
