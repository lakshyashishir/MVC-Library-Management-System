package views

import (
	"html/template"
)

func AdminPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/admin.html"))
	return temp
}

func AdminRequestsPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/adminRequests.html"))
	return temp
}

func IssuedBooksPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/issuedBooks.html"))
	return temp
}

func AddBookPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/addBook.html"))
	return temp
}

func RequestsPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/requests.html"))
	return temp
}
