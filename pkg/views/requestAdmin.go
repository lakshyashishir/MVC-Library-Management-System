package views

import (
	"html/template"
)

func RequestAdminPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/pendingAdminApproval.html"))
	return temp
}
