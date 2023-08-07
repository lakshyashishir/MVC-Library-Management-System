package views

import (
	"html/template"
)

func UserPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/user.html"))
	return temp
}

func UserRequestsPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/myrequests.html"))
	return temp
}

func UserBooksPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/mybooks.html"))
	return temp
}

func UserViewBookPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/getBook.html"))
	return temp
}
