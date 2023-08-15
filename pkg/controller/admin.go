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
	getUser, err := models.Auth(w, r)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username := getUser.Username
	t := views.AdminPage()
	t.Execute(w, username)
}

func AdminRequests(w http.ResponseWriter, r *http.Request) {
	CheckRoleAdmin(w, r)
	t := views.AdminRequestsPage()
	requests, err := models.GetAdminRequests()
	if err != nil {
		http.Error(w, "Error getting requests", http.StatusInternalServerError)
		return
	}
	t.Execute(w, requests)
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
	// CheckRoleAdmin(w, r)
	pendingRequests, err := models.GetAllPendingRequests()
	if err != nil {
		http.Error(w, "Error getting requests", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	rejectedRequests, err := models.GetAllRejectedRequests()
	if err != nil {
		http.Error(w, "Error getting requests", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	data := struct {
		PendingRequests  []types.RequestAdminView
		RejectedRequests []types.RequestAdminView
	}{
		PendingRequests:  pendingRequests,
		RejectedRequests: rejectedRequests,
	}

	t := views.RequestsPage()
	t.Execute(w, data)
}

func ApproveAdmin(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userID")
	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting userID to integer: %s", err), http.StatusInternalServerError)
		return
	}

	CheckRoleAdmin(w, r)
	err = models.ApproveAdminPost(userIDInt)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error approving admin: %s", err), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func ApproveBookRequest(w http.ResponseWriter, r *http.Request) {
	requestIDStr := r.FormValue("requestID")
	bookIdStr := r.FormValue("bookId")

	requestID, err := strconv.Atoi(requestIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting requestID to integer: %s", err), http.StatusInternalServerError)
		return
	}

	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting bookId to integer: %s", err), http.StatusInternalServerError)
		return
	}

	CheckRoleAdmin(w, r)
	err = models.ApproveBookRequestPost(requestID, bookId)
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

func AdminViewBook(w http.ResponseWriter, r *http.Request) {
	t := views.AdminViewBookPage()
	b, err := models.GetBook()
	if err != nil {
		http.Error(w, "Error getting book", http.StatusInternalServerError)
		return
	}
	// fmt.Println(b)
	t.Execute(w, b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting bookId to integer: %s", err), http.StatusInternalServerError)
		return
	}

	err = models.DeleteBookPost(bookIdInt)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting book: %s", err), http.StatusInternalServerError)
		return
	}

	CheckRoleAdmin(w, r)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting book: %s", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
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
