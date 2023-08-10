package controller

import (
	"fmt"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, request *http.Request) {
	fmt.Println("Login GET")
	t := views.LoginPage()
	t.Execute(w, nil)
}

func LoginPost(w http.ResponseWriter, request *http.Request) {
	fmt.Println("Login POST")
	username := request.FormValue("username")
	password := request.FormValue("password")

	fmt.Println(username, password)

	userRole, err := models.LoginPost(username, password)
	if err != nil {
		fmt.Println(err)
		return
	}

	sessionID, err := models.GeneratSessionID()
	if err != nil {
		fmt.Println(err)
		return
	}

	cookie := http.Cookie{
		Name:     "SessionID",
		Value:    sessionID,
		Expires:  time.Now().Add(48 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	fmt.Println("SessionID: ", sessionID)

	UserID, err := models.GetUserIDFromUsername(username)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = models.UpdateSessionID(UserID, sessionID)
	if err != nil {
		fmt.Println(err)
		return
	}

	if userRole == "admin" {
		http.Redirect(w, request, "/admin", http.StatusSeeOther)
	} else if userRole == "user" {
		http.Redirect(w, request, "/user", http.StatusSeeOther)
	} else if userRole == "admin requested" {
		http.Redirect(w, request, "/reqAdmin", http.StatusSeeOther)
	} else {
		http.Redirect(w, request, "/login", http.StatusSeeOther)
	}
}

func Logout(w http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("SessionID")
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, request, "/login", http.StatusSeeOther)
		return
	}
	fmt.Println(cookie.Value)
	err = models.DeleteSessionID(cookie.Value)
	if err != nil {
		fmt.Println(err)
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, request, "/login", http.StatusSeeOther)
}
