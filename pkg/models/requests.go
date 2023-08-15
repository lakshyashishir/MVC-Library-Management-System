package models

import (
	"fmt"
	"mvc/pkg/types"
)

func ApproveBookRequestPost(requestID int, bookID int) error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	rows, err := db.Query("Select * from books where book_id = ? AND status = 'available'", bookID)
	if err != nil {
		err = fmt.Errorf("error checking book status: %s", err)
		return err
	}
	if rows == nil {
		fmt.Println("Book is not available")
		return nil
	}

	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookID, &book.Title, &book.Author, &book.BookStatus, &book.Quantity)
		if err != nil {
			return fmt.Errorf("error scanning book rows: %s", err)
		}

		_, err = db.Query("UPDATE requests SET book_status = 'approved' WHERE book_id = ? and request_id = ?", bookID, requestID)
		if err != nil {
			return fmt.Errorf("error updating request status: %s", err)
		}

		if book.Quantity == 1 {
			_, err = db.Query("UPDATE books SET status = 'not available' WHERE book_id = ?", bookID)
			if err != nil {
				return fmt.Errorf("error updating book status: %s", err)
			}
		} else {
			_, err = db.Query("UPDATE books SET quantity = quantity - 1 WHERE book_id = ?", bookID)
			if err != nil {
				return fmt.Errorf("error updating book quantity: %s", err)
			}
		}

	}

	return nil
}

func RejectBookRequestPost(requestID int) error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	_, err = db.Query("UPDATE requests SET book_status = 'rejected' WHERE request_id = ?", requestID)
	if err != nil {
		return fmt.Errorf("error deleting book: %s", err)
	}

	return nil
}

func GetAllPendingRequests() ([]types.RequestAdminView, error) {
	db, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var RequestList []types.RequestAdminView

	rows, err := db.Query("SELECT * FROM requests WHERE book_status = 'pending'")
	if err != nil {
		return nil, fmt.Errorf("error searching requests: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var request types.RequestAdminView
		err := rows.Scan(&request.RequestID, &request.UserID, &request.BookID, &request.RequestStatus)
		if err != nil {
			return nil, fmt.Errorf("error scanning request rows: %s", err)
		}

		// fmt.Println(request.BookID)

		request.Username, err = GetUsernameFromID(request.UserID)
		if err != nil {
			return nil, fmt.Errorf("error getting username: %s", err)
		}
		request.Title, err = GetBookTitleByBookID(request.BookID)
		if err != nil {
			return nil, fmt.Errorf("error getting book title: %s", err)
		}
		request.BookStatus, err = GetBookStatusByBookID(request.BookID)
		if err != nil {
			return nil, fmt.Errorf("error getting book author: %s", err)
		}
		RequestList = append(RequestList, request)
	}

	return RequestList, nil
}

func GetAllRejectedRequests() ([]types.RequestAdminView, error) {
	db, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var RequestList []types.RequestAdminView

	rows, err := db.Query("SELECT * FROM requests WHERE book_status = 'rejected'")
	if err != nil {
		return nil, fmt.Errorf("error searching requests: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var request types.RequestAdminView
		err := rows.Scan(&request.RequestID, &request.UserID, &request.BookID, &request.RequestStatus)
		if err != nil {
			return nil, fmt.Errorf("error scanning request rows: %s", err)
		}

		// fmt.Println(request.BookID)

		request.Username, err = GetUsernameFromID(request.UserID)
		if err != nil {
			return nil, fmt.Errorf("error getting username: %s", err)
		}
		request.Title, err = GetBookTitleByBookID(request.BookID)
		if err != nil {
			return nil, fmt.Errorf("error getting book title: %s", err)
		}
		request.BookStatus, err = GetBookStatusByBookID(request.BookID)
		if err != nil {
			return nil, fmt.Errorf("error getting book author: %s", err)
		}
		RequestList = append(RequestList, request)
	}

	return RequestList, nil
}
