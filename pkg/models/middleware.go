package models

import (
	"fmt"
	"mvc/pkg/types"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) (types.User, error) {
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println("not authenticated")
		return types.User{}, err
	}

	db, err := Connect()
	if err != nil {
		fmt.Println(err)
		return types.User{}, err
	}

	defer db.Close()

	var user types.User

	row := db.QueryRow("SELECT cookies.userId FROM cookies WHERE sessionId = ?", cookie.Value)
	var userID int
	if err := row.Scan(&userID); err != nil {
		fmt.Println(err)
		return types.User{}, err
	}

	row = db.QueryRow("SELECT * FROM users WHERE  user_id = ?", userID)
	if err := row.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Role); err != nil {
		fmt.Println(err)
		return types.User{}, err
	}

	return user, nil
}
