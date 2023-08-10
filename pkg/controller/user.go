package controller

import (
	"fmt"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {
	getUser, err := models.Auth(w, r)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username := getUser.Username
	CheckRoleUser(w, r)

	fmt.Println(username)
	t := views.UserPage()
	t.Execute(w, nil)
}

func UserRequests(w http.ResponseWriter, request *http.Request) {
	CheckRoleUser(w, request)
	t := views.UserRequestsPage()
	t.Execute(w, nil)
}

func UserBooks(w http.ResponseWriter, request *http.Request) {
	CheckRoleUser(w, request)
	t := views.UserBooksPage()
	t.Execute(w, nil)
}

func UserViewBook(w http.ResponseWriter, request *http.Request) {
	CheckRoleUser(w, request)
	t := views.UserViewBookPage()
	b, err := models.GetBook()
	if err != nil {
		http.Error(w, "Error getting book", http.StatusInternalServerError)
		return
	}
	fmt.Println(b)
	t.Execute(w, b)
}

func CheckRoleUser(w http.ResponseWriter, r *http.Request) {
	getUser, err := models.Auth(w, r)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	userRole := getUser.Role
	if userRole == "admin" {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
	if userRole == "admin requested" {
		http.Redirect(w, r, "/reqAdmin", http.StatusSeeOther)
	}
}
