package views

import (
	"html/template"
)

func LoginPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/login.html"))
	return temp
}

func LoginAdminPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/loginAdmin.html"))
	return temp
}
