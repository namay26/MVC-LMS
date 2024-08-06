package views

import (
	"html/template"
	"io"

	"github.com/namay26/MVC-LMS/pkg/types"
)

func Render(w io.Writer, name string, data types.Datasent) {
	tmpl, err := template.ParseFiles("pkg/views/templates/" + name + ".html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, data)
}
