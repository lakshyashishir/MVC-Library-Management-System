package controller

import (
	"log"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, request *http.Request) {
	t := views.LoginPage()
	t.Execute(w, nil)
}

func LoginAdmin(w http.ResponseWriter, request *http.Request) {
	t := views.LoginAdminPage()
	t.Execute(w, nil)
}

func LoginPost(w http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")
	password := request.FormValue("password")

	userRole, err := models.LoginPost(username, password)
	if err != nil {
		log.Println(err)
		http.Redirect(w, request, "/500", http.StatusSeeOther)
		return
	}

	sessionID, err := models.GeneratSessionID()
	if err != nil {
		log.Println(err)
		http.Redirect(w, request, "/500", http.StatusSeeOther)
		return
	}

	cookie := http.Cookie{
		Name:     "sessionId",
		Value:    sessionID,
		Expires:  time.Now().Add(48 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	UserID, err := models.GetUserIDFromUsername(username)
	if err != nil {
		log.Println(err)
		http.Redirect(w, request, "/500", http.StatusSeeOther)
		return
	}

	err = models.UpdateSessionID(UserID, sessionID)
	if err != nil {
		log.Println(err)
		http.Redirect(w, request, "/500", http.StatusSeeOther)
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
	cookie, err := request.Cookie("sessionId")
	if err != nil {
		log.Println(err)
		http.Redirect(w, request, "/login", http.StatusSeeOther)
		return
	}
	err = models.DeleteSessionID(cookie.Value)
	if err != nil {
		log.Println(err)
		http.Redirect(w, request, "/login", http.StatusSeeOther)
		return
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, request, "/login", http.StatusSeeOther)
}
