package controller

import (
	"log"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		t := views.SignupPage()
		t.Execute(w, nil)
	case "POST":
		username := request.FormValue("username")
		password := request.FormValue("password")
		role := request.FormValue("role")

		if CheckUsernameExist(username) {
			log.Printf("Username already exist")
			return
		}

		adminExists := CheckAdminExist()

		if role == "admin" && adminExists {
			role = "admin requested"
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			http.Redirect(w, request, "/500", http.StatusSeeOther)
			return
		}

		salt := "salt"

		user := types.User{
			Username: username,
			Hash:     string(hash),
			Salt:     salt,
			Role:     types.UserRole(role),
		}

		errSignup := models.SignupPost(user)
		if errSignup != nil {
			log.Println(errSignup)
			http.Redirect(w, request, "/500", http.StatusSeeOther)
			return
		}

		sessionId, err := models.GeneratSessionID()
		if err != nil {
			log.Println(err)
			return
		}

		cookie := http.Cookie{
			Name:     "sessionId",
			Value:    sessionId,
			Expires:  time.Now().Add(48 * time.Hour),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, &cookie)

		userId, err := models.GetUserIDFromUsername(username)
		if err != nil {
			log.Println(err)
			http.Redirect(w, request, "/500", http.StatusSeeOther)
			return
		}

		err = models.UpdateSessionID(userId, sessionId)
		if err != nil {
			log.Println(err)
			http.Redirect(w, request, "/500", http.StatusSeeOther)
			return
		}

		if role == "admin" {
			http.Redirect(w, request, "/admin", http.StatusSeeOther)
		} else if role == "user" {
			http.Redirect(w, request, "/user", http.StatusSeeOther)
		} else if role == "admin requested" {
			http.Redirect(w, request, "/pendingAdminApproval", http.StatusSeeOther)
		} else {
			http.Redirect(w, request, "/login", http.StatusSeeOther)
		}

		return
	}
}

func CheckUsernameExist(username string) bool {
	db, err := models.Connect()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	return rows.Next()
}

func CheckAdminExist() bool {
	db, err := models.Connect()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE role = 'admin'")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	return rows.Next()
}
