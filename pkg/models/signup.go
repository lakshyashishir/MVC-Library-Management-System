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
	UserList := User

	// fmt.Println(UserList)

	_, err = db.Query("INSERT INTO users (username, hash, salt, role) VALUES (?, ?, ?, ?)", UserList.Username, UserList.Hash, UserList.Salt, UserList.Role)
	if err != nil {
		return fmt.Errorf("error registering User: %s", err)
	}

	// fmt.Println("User registered successfully")

	return nil
}
