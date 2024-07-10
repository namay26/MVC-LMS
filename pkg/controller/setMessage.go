package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/namay26/MVC-LMS/pkg/model"
)

var SecretKey string

func init() {
	SecretKey = model.JwtSecretKey()
}

func SetFlash(writer http.ResponseWriter, request *http.Request, message string) {
	store := sessions.NewCookieStore([]byte(SecretKey))
	session, err := store.Get(request, "flash-session")
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
	session.AddFlash(message, "message")
	session.Save(request, writer)
}

func GetFlash(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
	store := sessions.NewCookieStore([]byte(SecretKey))
	session, err := store.Get(request, "flash-session")
	if err != nil {
		return "", err
	}
	fm := session.Flashes("message")
	if fm == nil {
		return "", nil
	}
	session.Save(request, writer)
	message := fm[0]
	return message, err
}
