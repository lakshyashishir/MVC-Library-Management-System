package models

import (
	"fmt"
	"mvc/pkg/types"
)

func GetAdminRequests() ([]types.User, error) {
	db, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	// fmt.Println("checking")
	var UserList []types.User

	rows, err := db.Query("SELECT user_id,username, hash, salt, role FROM users WHERE role = 'admin requested'")
	if err != nil {
		return nil, fmt.Errorf("error searching users: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Role)
		if err != nil {
			return nil, fmt.Errorf("error searching users: %s", err)
		}
		// fmt.Println(user)
		UserList = append(UserList, user)
	}

	// fmt.Println(UserList)
	// fmt.Println("checking")

	return UserList, nil
}

func ApproveAdminPost(userID int) error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	_, err = db.Query("UPDATE users SET role = 'admin' WHERE user_id = ?", userID)
	if err != nil {
		return fmt.Errorf("error updating user: %s", err)
	}

	return nil
}

func RejectAdminPost(userID int) error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	_, err = db.Query("Delete from users WHERE user_id = ?", userID)
	if err != nil {
		return fmt.Errorf("error updating user: %s", err)
	}

	return nil
}
