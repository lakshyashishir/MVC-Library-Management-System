package controller

import (
	"log"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

func User(w http.ResponseWriter, r *http.Request) {
	getUser, err := models.Auth(w, r)
	if err != nil {
		log.Print(err)
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
		log.Print(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	userId := getUser.UserID
	bookId := r.FormValue("bookId")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	err = models.UserRequestBookPost(userId, bookIdInt)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func UserRequests(w http.ResponseWriter, r *http.Request) {
	CheckRoleUser(w, r)

	getUser, err := models.Auth(w, r)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	userId := getUser.UserID

	myPendingRequests, err := models.GetUserRequestsPending(userId)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	myRejectedRequests, err := models.GetUserRequestsRejected(userId)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	data := struct {
		PendingRequests  []types.BookUserView
		RejectedRequests []types.BookUserView
	}{
		PendingRequests:  myPendingRequests,
		RejectedRequests: myRejectedRequests,
	}

	t := views.UserRequestsPage()
	t.Execute(w, data)
}

func UserBooks(w http.ResponseWriter, r *http.Request) {
	CheckRoleUser(w, r)
	getUser, err := models.Auth(w, r)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	userId := getUser.UserID
	userBooks, err := models.GetUserBooks(userId)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	t := views.UserBooksPage()
	t.Execute(w, userBooks)
}

func UserViewBook(w http.ResponseWriter, r *http.Request) {
	t := views.UserViewBookPage()
	b, err := models.GetBook()
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	t.Execute(w, b)
}

func UserRemoveRequestBook(w http.ResponseWriter, r *http.Request) {
	requestId := r.FormValue("requestId")

	requestIdInt, err := strconv.Atoi(requestId)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	bookId, err := models.GetBookIdByRequestId(requestIdInt)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	models.UserRemoveRequestPost(requestIdInt, bookId)

	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func UserReturnBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")

	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	userIdInt, err := models.GetUserIdByBookId(bookIdInt)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	models.UserReturnBookPost(bookIdInt, userIdInt)

	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func CheckRoleUser(w http.ResponseWriter, r *http.Request) {
	getUser, err := models.Auth(w, r)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	userRole := getUser.Role
	if userRole == "admin" {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
	if userRole == "admin requested" {
		http.Redirect(w, r, "/pendingAdminApproval", http.StatusSeeOther)
	}
}
