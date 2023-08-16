package controller

import (
	"log"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

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
