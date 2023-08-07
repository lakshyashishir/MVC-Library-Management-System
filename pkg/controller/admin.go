package controller

import (
	"mvc/pkg/views"
	"net/http"
)

func Admin(w http.ResponseWriter, request *http.Request) {
	t := views.AdminPage()
	t.Execute(w, nil)
}

func AdminRequests(w http.ResponseWriter, request *http.Request) {
	t := views.AdminRequestsPage()
	t.Execute(w, nil)
}

func IssuedBooks(w http.ResponseWriter, request *http.Request) {
	t := views.IssuedBooksPage()
	t.Execute(w, nil)
}

func AddBook(w http.ResponseWriter, request *http.Request) {
	t := views.AddBookPage()
	t.Execute(w, nil)
}

func Requests(w http.ResponseWriter, request *http.Request) {
	t := views.RequestsPage()
	t.Execute(w, nil)
}
