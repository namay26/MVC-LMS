package controller

import (
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/middleware"
	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func GetReqCheckout(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()
	user := middleware.GetUser(r)
	booklist, err := model.RequestCheckout(db, user)
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

	user := middleware.GetUser(r)

	err := model.Checkout(db, bookid, user.Userid)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/user/reqcheckout", http.StatusFound)

}
