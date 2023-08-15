package controller

import (
	"fmt"
	"mvc/pkg/models"
	"mvc/pkg/types"
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

	t := views.UserPage()
	t.Execute(w, username)
}

func UserRequestBook(w http.ResponseWriter, r *http.Request) {
	CheckRoleUser(w, r)
	getUser, err := models.Auth(w, r)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	userId := getUser.UserID
	bookId := r.FormValue("bookId")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = models.UserRequestBookPost(userId, bookIdInt)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func UserRequests(w http.ResponseWriter, r *http.Request) {
	CheckRoleUser(w, r)

	getUser, err := models.Auth(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	userId := getUser.UserID

	MyPendingRequests, err := models.GetUserRequestsPending(userId)
	if err != nil {
		fmt.Println(err)
		return
	}

	MyRejectedRequests, err := models.GetUserRequestsRejected(userId)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		PendingRequests  []types.BookUserView
		RejectedRequests []types.BookUserView
	}{
		PendingRequests:  MyPendingRequests,
		RejectedRequests: MyRejectedRequests,
	}

	t := views.UserRequestsPage()
	t.Execute(w, data)
}

func UserBooks(w http.ResponseWriter, r *http.Request) {
	CheckRoleUser(w, r)
	getUser, err := models.Auth(w, r)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	userId := getUser.UserID
	UserBooks, err := models.GetUserBooks(userId)
	if err != nil {
		fmt.Println(err)
		return
	}

	t := views.UserBooksPage()
	t.Execute(w, UserBooks)
}

func UserViewBook(w http.ResponseWriter, r *http.Request) {
	t := views.UserViewBookPage()
	b, err := models.GetBook()
	if err != nil {
		http.Error(w, "Error getting book", http.StatusInternalServerError)
		return
	}
	// fmt.Println(b)
	t.Execute(w, b)
}

func UserRemoveRequestBook(w http.ResponseWriter, r *http.Request) {
	requestID := r.FormValue("requestID")

	requestIDInt, err := strconv.Atoi(requestID)
	if err != nil {
		fmt.Println(err)
		return
	}
	bookId, err := models.GetBookIdByRequestId(requestIDInt)
	if err != nil {
		fmt.Println(err)
		return
	}

	models.UserRemoveRequestPost(requestIDInt, bookId)

	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func UserReturnBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")

	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println(err)
		return
	}

	userIdInt, err := models.GetUserIdByBookId(bookIdInt)
	if err != nil {
		fmt.Println(err)
		return
	}
	models.UserReturnBookPost(bookIdInt, userIdInt)

	http.Redirect(w, r, "/user", http.StatusSeeOther)
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
