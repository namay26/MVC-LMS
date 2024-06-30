package main

import (
	"fmt"
	"net/http"

	"github.com/namay26/MVC-LMS/controller"
)

func main() {
	http.HandleFunc("/", controller.GetPage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
