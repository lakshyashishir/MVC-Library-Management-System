package controller

import (
	"fmt"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
)

func RequestAdmin(w http.ResponseWriter, r *http.Request) {
	CheckRoleAdminRequest(w, r)
	t := views.RequestAdminPage()
	t.Execute(w, nil)
}

func CheckRoleAdminRequest(w http.ResponseWriter, r *http.Request) {
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
	if userRole == "admin" {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
