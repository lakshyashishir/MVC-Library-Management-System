package models

import (
	"database/sql"
	"fmt"
	"mvc/pkg/types"
)

func GetAdminRequests(db *sql.DB) ([]types.User, error) {
	var userList []types.User

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
		userList = append(userList, user)
	}
	return userList, nil
}

func ApproveAdminPost(userId int) error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	_, err = db.Query("UPDATE users SET role = 'admin' WHERE user_id = ?", userId)
	if err != nil {
		return fmt.Errorf("error updating user: %s", err)
	}

	return nil
}

func RejectAdminPost(userId int) error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	_, err = db.Query("Delete from users WHERE user_id = ?", userId)
	if err != nil {
		return fmt.Errorf("error updating user: %s", err)
	}

	return nil
}
