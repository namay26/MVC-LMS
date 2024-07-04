package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/namay26/MVC-LMS/controller"
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

	mainRouter.HandleFunc("/login", controller.Login).Methods("POST")
	mainRouter.HandleFunc("/register", controller.RegisterPage).Methods("POST")

	userRouter.HandleFunc("/home", controller.Home).Methods("GET")
	userRouter.HandleFunc("/listbooks", controller.ListBooks).Methods("GET")

	adminRouter.HandleFunc("/home", controller.AdminHome).Methods("GET")
	adminRouter.HandleFunc("/listbooks", controller.AdminListBook).Methods("GET")
	adminRouter.HandleFunc("/listbooks", controller.AdminUpdateCheck).Methods("POST")

	adminRouter.HandleFunc("/updatebook", controller.AdminUpdate).Methods("POST")

	adminRouter.HandleFunc("/addbook", controller.GetAddBook).Methods("GET")
	adminRouter.HandleFunc("/addbook", controller.AddBook).Methods("POST")
	// router.HandleFunc("/500", controller.InternalServerError).Methods("GET")
	// router.HandleFunc("/403", controller.UnauthorizedAccessError).Methods("GET")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
