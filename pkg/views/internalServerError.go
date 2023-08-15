package views

import (
	"html/template"
)

func InternalServerErrorPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/internalServerError.html"))
	return temp
}
