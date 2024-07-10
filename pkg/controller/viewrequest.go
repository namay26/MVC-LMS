package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetViewRequest(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	listusers, _ := model.ViewRequest(db)
	views.Render(w, "viewrequest", listusers)
}

func ViewRequest(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	r.ParseForm()
	userid := r.FormValue("userid")
	bookid := r.FormValue("bookid")
	acptreq, err := model.AcceptRequest(db, userid, bookid)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	if acptreq {
		http.Redirect(w, r, "/admin/viewrequest", http.StatusSeeOther)
	} else {
		fmt.Println("Error")
	}

}
