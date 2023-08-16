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
	db, err := models.Connect()
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Println(err)
		return
	}
	defer db.Close()
	requests, err := models.GetAdminRequests(db)
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		log.Println(err)
		return
	}
	t.Execute(w, requests)
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
		http.Redirect(w, r, "/pendingAdminApproval", http.StatusSeeOther)
	}
}
