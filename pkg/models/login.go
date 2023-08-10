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

	var User types.User
	err = db.QueryRow("SELECT user_id, hash, role FROM users WHERE username = ?", username).Scan(&User.UserID, &User.Hash, &User.Role)
	if err != nil {
		return "", fmt.Errorf("error searching user: %s", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(User.Hash), []byte(password))
	if err != nil {
		return "", fmt.Errorf("error comparing password: %s", err)
	}

	return User.Role, nil
}

func GetUserIDFromUsername(username string) (int, error) {
	db, err := Connect()
	if err != nil {
		return 0, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var UserID int
	err = db.QueryRow("SELECT user_id FROM users WHERE username = ?", username).Scan(&UserID)
	if err != nil {
		return 0, fmt.Errorf("error searching user: %s", err)
	}

	return UserID, nil
}
