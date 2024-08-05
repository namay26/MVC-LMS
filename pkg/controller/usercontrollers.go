package controller

import (
	"log"
	"net/http"

	"github.com/namay26/MVC-LMS/pkg/middleware"
	"github.com/namay26/MVC-LMS/pkg/model"
	"github.com/namay26/MVC-LMS/pkg/structs"
	"github.com/namay26/MVC-LMS/pkg/views"
)

func UserHome(w http.ResponseWriter, r *http.Request) {
	var data structs.Datasent
	views.Render(w, "home", data)
}

// Check User Borrow-history

func GetBorrowHistory(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()
	user := middleware.GetUser(r)
	borrowhistory, err := model.GetBorrowHistory(db, user)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var data structs.Datasent
	data.Results = borrowhistory
	views.Render(w, "borrowhistory", data)
}

// User Book List

func UserListBooks(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	booklist, err := model.GetBooks(db)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var data structs.Datasent
	data.Results = booklist
	views.Render(w, "listbooks", data)
}

// User Book Checkout Request

func GetReqCheckout(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()
	user := middleware.GetUser(r)
	booklist, err2 := model.RequestCheckout(db, user)
	if err2 != nil {
		panic(err2)
	}
	var message structs.PageMessage
	var err error
	message.Message, err = GetFlash(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var data structs.Datasent
	data.Message = message
	data.Results = booklist
	views.Render(w, "reqcheckout", data)
}

func ReqCheckout(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	r.ParseForm()
	bookid := r.FormValue("bookid")

	user := middleware.GetUser(r)

	err := model.Checkout(db, bookid, user.Userid)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	SetFlash(w, r, "Checkout Requested Successfully!")
	http.Redirect(w, r, "/user/reqcheckout", http.StatusFound)
}

// User Request For Admin Privileges

func GetRequestAdmin(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	user := middleware.GetUser(r)

	check, err1 := model.CheckRequest(db, user)

	if err1 != nil {
		log.Println(err1)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	if check {
		SetFlash(w, r, "Request already sent")
	}

	var message structs.PageMessage
	var err error
	message.Message, err = GetFlash(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var data structs.Datasent
	data.Message = message
	views.Render(w, "requestadmin", data)
}

func RequestAdmin(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	user := middleware.GetUser(r)

	check, err1 := model.CheckRequest(db, user)

	if err1 != nil {
		log.Println(err1)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	if check {
		SetFlash(w, r, "Request already sent")
		return
	}
	requestadmin, err := model.RequestAdmin(db, user)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	if !requestadmin {
		SetFlash(w, r, "Couldn't make the request")
		http.Redirect(w, r, "/user/requestadmin", http.StatusSeeOther)
		return
	} else {
		SetFlash(w, r, "Request sent successfully")
		http.Redirect(w, r, "/user/requestadmin", http.StatusSeeOther)
		return
	}
}

// User Return Book

func GetReturnBook(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	user := middleware.GetUser(r)

	returnbook, _ := model.GetReturnBook(db, user)
	var message structs.PageMessage
	var err error
	message.Message, err = GetFlash(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	var data structs.Datasent
	data.Message = message
	data.Results = returnbook
	views.Render(w, "returnbook", data)

}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	db, _ := model.Connect()
	defer db.Close()

	bookid := r.FormValue("bookid")

	user := middleware.GetUser(r)

	err := model.ReturnBook(db, user, bookid)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	SetFlash(w, r, "Book Returned Successfully!")
	http.Redirect(w, r, "/user/returnbook", http.StatusSeeOther)
}
