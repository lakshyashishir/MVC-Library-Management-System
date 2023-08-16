package models

import (
	"fmt"
	"mvc/pkg/types"
)

func SignupPost(User types.User) error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()
	userList := User

	_, err = db.Query("INSERT INTO users (username, hash, salt, role) VALUES (?, ?, ?, ?)", userList.Username, userList.Hash, userList.Salt, userList.Role)
	if err != nil {
		return fmt.Errorf("error registering User: %s", err)
	}

	return nil
}
