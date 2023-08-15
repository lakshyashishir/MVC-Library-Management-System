package views

import (
	"html/template"
)

func NotFoundPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/404.html"))
	return temp
}
