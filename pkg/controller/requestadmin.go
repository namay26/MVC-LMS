package controller

import (
	"log"
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/middleware"
	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/structs"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetRequestAdmin(w http.ResponseWriter, r *http.Request) {
	var message structs.PageMessage
	var err error
	message.Message, err = GetFlash(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	views.Render(w, "requestadmin", message)
}

func RequestAdmin(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	user := middleware.GetUser(r)

	requestadmin, err := model.RequestAdmin(db, user)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	if !requestadmin {
		SetFlash(w, r, "Couldn't make the request")
		views.Render(w, "requestadmin", "Request failed")
		return
	} else {
		SetFlash(w, r, "Request sent successfully")
		views.Render(w, "requestadmin", "Request sent")
		return
	}
}
