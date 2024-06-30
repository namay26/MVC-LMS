package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/views"
)

func GetPage(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "index", nil)
}
