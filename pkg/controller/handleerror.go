package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/views"
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "InternalServerError", nil)
}
