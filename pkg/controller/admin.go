package controller

import (
	"log"
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
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Println(err)
		return
	}
	t.Execute(w, requests)
}

func IssuedBooks(w http.ResponseWriter, r *http.Request) {
	CheckRoleAdmin(w, r)
	t := views.IssuedBooksPage()
	requests, err := models.GetIssuedBooks()
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Println(err)
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
			http.Redirect(w, r, "/500", http.StatusSeeOther)
			log.Printf("Error converting quantity to integer: %s", err)
			return
		}

		book := types.Book{
			Title:    title,
			Author:   author,
			Quantity: quantity,
		}

		err = models.AddBookPost(book)
		if err != nil {
			http.Redirect(w, r, "/500", http.StatusSeeOther)
			log.Printf("Error adding book: %s", err)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

func Requests(w http.ResponseWriter, r *http.Request) {
	CheckRoleAdmin(w, r)
	pendingRequests, err := models.GetAllPendingRequests()
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Println(err)
		return
	}

	rejectedRequests, err := models.GetAllRejectedRequests()
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Println(err)
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
	userId := r.FormValue("userId")
	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Printf("Error converting userId to integer: %s", err)
		return
	}

	CheckRoleAdmin(w, r)
	err = models.ApproveAdminPost(userIdInt)
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Printf("Error approving admin: %s", err)
		return
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func RejectAdmin(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("userId")
	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Printf("Error converting userId to integer: %s", err)
		return
	}

	CheckRoleAdmin(w, r)
	err = models.RejectAdminPost(userIdInt)
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Printf("Error rejecting admin: %s", err)
		return
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func ApproveBookRequest(w http.ResponseWriter, r *http.Request) {
	requestIdStr := r.FormValue("requestId")

	requestId, err := strconv.Atoi(requestIdStr)
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Printf("Error converting requestId to integer: %s", err)
		return
	}

	bookId, err := models.GetBookIdByRequestId(requestId)
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Printf("Error getting bookId: %s", err)
		return
	}

	CheckRoleAdmin(w, r)
	err = models.ApproveBookRequestPost(requestId, bookId)
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Printf("Error approving book request: %s", err)
		return
	}

	http.Redirect(w, r, "/requests", http.StatusSeeOther)
}

func RejectBookRequest(w http.ResponseWriter, r *http.Request) {
	requestIdStr := r.FormValue("requestId")

	requestId, err := strconv.Atoi(requestIdStr)
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Printf("Error converting requestId to integer: %s", err)
		return
	}

	CheckRoleAdmin(w, r)
	err = models.RejectBookRequestPost(requestId)
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Printf("Error rejecting book request: %s", err)
		return
	}

	http.Redirect(w, r, "/requests", http.StatusSeeOther)
}

func AdminViewBook(w http.ResponseWriter, r *http.Request) {
	t := views.AdminViewBookPage()
	b, err := models.GetBook()
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Printf("Error getting book")
		return
	}
	// fmt.Println(b)
	t.Execute(w, b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Printf("Error converting bookId to integer: %s", err)
		return
	}

	err = models.DeleteBookPost(bookIdInt)

	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Printf("Error deleting book: %s", err)
		return
	}

	CheckRoleAdmin(w, r)

	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Printf("Error deleting book: %s", err)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func CheckRoleAdmin(w http.ResponseWriter, r *http.Request) {
	getUser, err := models.Auth(w, r)
	if err != nil {
		log.Println(err)
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
