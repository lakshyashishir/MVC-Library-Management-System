package controller

import (
	"fmt"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

func Admin(w http.ResponseWriter, request *http.Request) {
	t := views.AdminPage()
	t.Execute(w, nil)
}

func AdminRequests(w http.ResponseWriter, request *http.Request) {
	t := views.AdminRequestsPage()
	requests, err := models.GetAdminRequests()
	fmt.Println(requests)
	if err != nil {
		http.Error(w, "Error getting requests", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func IssuedBooks(w http.ResponseWriter, request *http.Request) {
	t := views.IssuedBooksPage()
	requests, err := models.GetIssuedBooks()
	if err != nil {
		http.Error(w, "Error getting issued books", http.StatusInternalServerError)
		return
	}
	t.Execute(w, requests)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
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

func Requests(w http.ResponseWriter, request *http.Request) {
	t := views.RequestsPage()
	t.Execute(w, nil)
}

func ApproveAdmin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	err := models.ApproveAdminPost(username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error approving admin: %s", err), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
