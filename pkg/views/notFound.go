package views

import (
	"html/template"
)

func NotFoundPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/notFound.html"))
	return temp
}
