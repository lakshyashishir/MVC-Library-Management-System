package models

import (
	"fmt"
	"mvc/pkg/types"
)

func GetIssuedBooks() ([]types.Request, error) {
	db, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	Issuedlist := []types.Request{}

	rows, err := db.Query("SELECT * FROM requests where book_status = 'approved'")
	if err != nil {
		return nil, fmt.Errorf("error querying DB: %s", err)
	}

	for rows.Next() {
		var requests types.Request
		err := rows.Scan(&requests.RequestID, &requests.UserID, &requests.BookID, &requests.BookStatus)
		if err != nil {
			return nil, fmt.Errorf("error scanning rows: %s", err)
		}
		Issuedlist = append(Issuedlist, requests)
	}

	// log.Println(Issuedlist)

	return Issuedlist, nil
}
