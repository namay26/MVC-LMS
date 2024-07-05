package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/namay26/MVC-LMS/pkg/controller"
	"github.com/namay26/MVC-LMS/pkg/middleware"
)

func Initialize() {
	fmt.Println("Welcome to Library Manager")
	fmt.Println("Your server has started on port 8080")

	mainRouter := mux.NewRouter()
	mainRouter.HandleFunc("/", controller.GetPage)

	adminRouter := mainRouter.PathPrefix("/admin").Subrouter()
	userRouter := mainRouter.PathPrefix("/user").Subrouter()
	http.Handle("/", mainRouter)
	mainRouter.HandleFunc("/register", controller.RegisterPage).Methods("GET")
	mainRouter.HandleFunc("/login", controller.LoginPage).Methods("GET")

	adminRouter.Use(middleware.Authenticator)
	userRouter.Use(middleware.Authenticator)

	mainRouter.HandleFunc("/login", controller.Login).Methods("POST")
	mainRouter.HandleFunc("/register", controller.Register).Methods("POST")

	userRouter.HandleFunc("/home", controller.Home).Methods("GET")
	userRouter.HandleFunc("/listbooks", controller.ListBooks).Methods("GET")

	adminRouter.HandleFunc("/home", controller.AdminHome).Methods("GET")
	adminRouter.HandleFunc("/listbooks", controller.AdminListBook).Methods("GET")
	adminRouter.HandleFunc("/listbooks", controller.AdminUpdateCheck).Methods("POST")

	adminRouter.HandleFunc("/updatebook", controller.AdminUpdate).Methods("POST")

	adminRouter.HandleFunc("/addbook", controller.GetAddBook).Methods("GET")
	adminRouter.HandleFunc("/addbook", controller.AddBook).Methods("POST")

	adminRouter.HandleFunc("/deletebook", controller.GetDeleteBook).Methods("GET")
	adminRouter.HandleFunc("/deletebook", controller.DeleteBook).Methods("POST")

	userRouter.HandleFunc("/reqcheckout", controller.GetReqCheckout).Methods("GET")
	userRouter.HandleFunc("/reqcheckout", controller.ReqCheckout).Methods("POST")

	userRouter.HandleFunc("/borrowhistory", controller.GetBorrowHistory).Methods("GET")

	userRouter.HandleFunc("/returnbook", controller.GetReturnBook).Methods("GET")
	userRouter.HandleFunc("/returnbook", controller.ReturnBook).Methods("POST")

	userRouter.HandleFunc("/requestadmin", controller.GetRequestAdmin).Methods("GET")
	userRouter.HandleFunc("/requestadmin", controller.RequestAdmin).Methods("POST")

	adminRouter.HandleFunc("/grantadmin", controller.GetGrantAdmin).Methods("GET")
	adminRouter.HandleFunc("/grantadmin", controller.GrantAdmin).Methods("POST")

	adminRouter.HandleFunc("/viewrequest", controller.GetViewRequest).Methods("GET")
	//adminRouter.HandleFunc("/viewrequest", controller.ViewRequest).Methods("POST")
	// router.HandleFunc("/500", controller.InternalServerError).Methods("GET")
	// router.HandleFunc("/403", controller.UnauthorizedAccessError).Methods("GET")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
