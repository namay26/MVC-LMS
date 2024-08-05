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

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	GetPage := http.HandlerFunc(controller.GetPage)
	LoginPage := http.HandlerFunc(controller.LoginPage)
	adminRouter := mainRouter.PathPrefix("/admin").Subrouter()
	userRouter := mainRouter.PathPrefix("/user").Subrouter()
	http.Handle("/", mainRouter)
	mainRouter.HandleFunc("/register", controller.RegisterPage).Methods("GET")
	mainRouter.HandleFunc("/logout", controller.Logout).Methods("GET")
	mainRouter.Handle("/", middleware.LoginMiddleware(GetPage)).Methods("GET")
	mainRouter.Handle("/login", middleware.LoginMiddleware(LoginPage)).Methods("GET")

	adminRouter.Use(middleware.Authenticator)
	userRouter.Use(middleware.Authenticator)

	mainRouter.HandleFunc("/login", controller.Login).Methods("POST")
	mainRouter.HandleFunc("/register", controller.Register).Methods("POST")

	userRouter.HandleFunc("/home", controller.UserHome).Methods("GET")
	userRouter.HandleFunc("/listbooks", controller.UserListBooks).Methods("GET")

	adminRouter.HandleFunc("/home", controller.AdminHome).Methods("GET")
	adminRouter.HandleFunc("/listbooks", controller.GetAdminListBook).Methods("GET")
	adminRouter.HandleFunc("/listbooks", controller.GetAdminUpdate).Methods("POST")

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
	adminRouter.HandleFunc("/viewrequest", controller.ViewRequest).Methods("POST")
	mainRouter.HandleFunc("/500", controller.InternalServerError).Methods("GET")

	mainRouter.NotFoundHandler = http.HandlerFunc(controller.PageNotFound)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
