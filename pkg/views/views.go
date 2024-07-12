package views

import (
	"html/template"
	"io"

	"github.com/namay26/MVC-LMS/pkg/structs"
)

func Render(w io.Writer, name string, data structs.Datasent) {
	tmpl, err := template.ParseFiles("pkg/views/templates/" + name + ".html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, data)
}
