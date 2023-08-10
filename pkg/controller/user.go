package controller

import (
	"fmt"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
	"strconv"
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

func UserRequests(w http.ResponseWriter, r *http.Request) {
	CheckRoleUser(w, r)
	switch r.Method {
	case "GET":
		t := views.UserRequestsPage()
		t.Execute(w, nil)
	case "POST":
		bookID := r.FormValue("bookID")
		getUser, err := models.Auth(w, r)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		UserID := getUser.UserID
		bookIDInt, err := strconv.Atoi(bookID)
		if err != nil {
			fmt.Println(err)
			return
		}
		models.UserRequestBookPost(bookIDInt, UserID)
	}
}

func UserBooks(w http.ResponseWriter, r *http.Request) {
	CheckRoleUser(w, r)
	t := views.UserBooksPage()
	t.Execute(w, nil)
}

func UserViewBook(w http.ResponseWriter, r *http.Request) {
	CheckRoleUser(w, r)
	t := views.UserViewBookPage()
	b, err := models.GetBook()
	if err != nil {
		http.Error(w, "Error getting book", http.StatusInternalServerError)
		return
	}
	fmt.Println(b)
	t.Execute(w, b)
}

func UserRemoveRequestBook(w http.ResponseWriter, r *http.Request) {
	requestID := r.FormValue("requestID")
	bookID := r.FormValue("bookID")

	requestIDInt, err := strconv.Atoi(requestID)
	if err != nil {
		fmt.Println(err)
		return
	}
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		fmt.Println(err)
		return
	}
	models.UserRemoveRequestPost(requestIDInt, bookIDInt)
}

func UserReturnBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("BookID")
	userID := r.FormValue("UserID")

	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		fmt.Println(err)
		return
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println(err)
		return
	}
	models.UserReturnBookPost(bookIDInt, userIDInt)
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
