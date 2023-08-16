package models

import (
	"fmt"
	"mvc/pkg/types"

	"golang.org/x/crypto/bcrypt"
)

func LoginPost(username string, password string) (types.UserRole, error) {
	db, err := Connect()
	if err != nil {
		return "", fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var user types.User
	err = db.QueryRow("SELECT user_id, hash, role FROM users WHERE username = ?", username).Scan(&user.UserID, &user.Hash, &user.Role)
	if err != nil {
		return "", fmt.Errorf("error searching user: %s", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
	if err != nil {
		return "", fmt.Errorf("error comparing password: %s", err)
	}

	return user.Role, nil
}
