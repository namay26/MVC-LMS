package views

import (
	"html/template"
	"io"
)

func Render(w io.Writer, name string, data interface{}) {
	tmpl, err := template.ParseFiles("../pkg/views/templates/" + name + ".html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, data)
}
