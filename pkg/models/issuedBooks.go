package models

import (
	"fmt"
	"mvc/pkg/types"
)

func GetIssuedBooks() ([]types.RequestAlt, error) {
	db, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	issuedList := []types.RequestAlt{}

	rows, err := db.Query("SELECT * FROM requests where book_status = 'approved'")
	if err != nil {
		return nil, fmt.Errorf("error querying DB: %s", err)
	}

	for rows.Next() {
		var requests types.RequestAlt
		err := rows.Scan(&requests.RequestID, &requests.UserID, &requests.BookID, &requests.BookStatus)
		if err != nil {
			return nil, fmt.Errorf("error scanning rows: %s", err)
		}
		requests.Username, err = GetUsernameFromID(requests.UserID)
		if err != nil {
			return nil, fmt.Errorf("error getting username: %s", err)
		}
		requests.Title, err = GetBookTitleByBookID(requests.BookID)
		if err != nil {
			return nil, fmt.Errorf("error getting book title: %s", err)
		}
		issuedList = append(issuedList, requests)
	}

	return issuedList, nil
}
