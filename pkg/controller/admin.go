package controller

import (
	"fmt"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	CheckRoleAdmin(w, r)
	t := views.AdminPage()
	t.Execute(w, nil)
}

func AdminRequests(w http.ResponseWriter, r *http.Request) {
	CheckRoleAdmin(w, r)
	t := views.AdminRequestsPage()
	requests, err := models.GetAdminRequests()
	fmt.Println(requests)
	if err != nil {
		http.Error(w, "Error getting requests", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func IssuedBooks(w http.ResponseWriter, r *http.Request) {
	CheckRoleAdmin(w, r)
	t := views.IssuedBooksPage()
	requests, err := models.GetIssuedBooks()
	if err != nil {
		http.Error(w, "Error getting issued books", http.StatusInternalServerError)
		return
	}
	t.Execute(w, requests)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	CheckRoleAdmin(w, r)
	switch r.Method {
	case "GET":
		t := views.AddBookPage()
		t.Execute(w, nil)

	case "POST":
		title := r.FormValue("title")
		author := r.FormValue("author")
		quantityStr := r.FormValue("quantity")
		quantity, err := strconv.Atoi(quantityStr)

		if err != nil {
			http.Error(w, fmt.Sprintf("Error converting quantity to integer: %s", err), http.StatusInternalServerError)
			return
		}

		book := types.Book{
			Title:    title,
			Author:   author,
			Quantity: quantity,
		}

		err = models.AddBookPost(book)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error adding book: %s", err), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

func Requests(w http.ResponseWriter, r *http.Request) {
	CheckRoleAdmin(w, r)
	t := views.RequestsPage()
	t.Execute(w, nil)
}

func ApproveAdmin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	CheckRoleAdmin(w, r)
	err := models.ApproveAdminPost(username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error approving admin: %s", err), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func ApproveBookRequest(w http.ResponseWriter, r *http.Request) {
	requestIDStr := r.FormValue("requestID")
	bookIDStr := r.FormValue("bookID")

	requestID, err := strconv.Atoi(requestIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting requestID to integer: %s", err), http.StatusInternalServerError)
		return
	}

	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting bookID to integer: %s", err), http.StatusInternalServerError)
		return
	}

	CheckRoleAdmin(w, r)
	err = models.ApproveBookRequestPost(requestID, bookID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error approving book request: %s", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/requests", http.StatusSeeOther)
}

func RejectBookRequest(w http.ResponseWriter, r *http.Request) {
	requestIDStr := r.FormValue("requestID")

	requestID, err := strconv.Atoi(requestIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting requestID to integer: %s", err), http.StatusInternalServerError)
		return
	}

	CheckRoleAdmin(w, r)
	err = models.RejectBookRequestPost(requestID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error rejecting book request: %s", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/requests", http.StatusSeeOther)
}

func CheckRoleAdmin(w http.ResponseWriter, r *http.Request) {
	getUser, err := models.Auth(w, r)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	userRole := getUser.Role
	if userRole == "user" {
		http.Redirect(w, r, "/user", http.StatusSeeOther)
	}
	if userRole == "admin requested" {
		http.Redirect(w, r, "/reqAdmin", http.StatusSeeOther)
	}
}
