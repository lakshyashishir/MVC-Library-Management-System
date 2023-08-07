package controller

import (
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
)

func User(w http.ResponseWriter, request *http.Request) {
	t := views.UserPage()
	t.Execute(w, nil)
}

func UserRequests(w http.ResponseWriter, request *http.Request) {
	t := views.UserRequestsPage()
	t.Execute(w, nil)
}

func UserBooks(w http.ResponseWriter, request *http.Request) {
	t := views.UserBooksPage()
	t.Execute(w, nil)
}

func UserViewBook(w http.ResponseWriter, request *http.Request) {
	t := views.UserViewBookPage()
	b := models.GetBook()
	t.Execute(w, b)
}
