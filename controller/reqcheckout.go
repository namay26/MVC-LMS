package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/model"
	"github.com/namay26/MVC-LMS/views"
)

func GetReqCheckout(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	booklist, err := model.RequestCheckout(db)
	if err != nil {
		panic(err)
	}
	views.Render(w, "reqcheckout", booklist)
}

func ReqCheckout(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	r.ParseForm()
	bookid := r.FormValue("bookid")
	userid := r.FormValue("userid")

	err := model.Checkout(db, bookid, userid)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/user/home", http.StatusFound)

}
