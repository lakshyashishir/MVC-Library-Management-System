package models

import (
	"fmt"
	"mvc/pkg/types"
)

func GetRequests() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	BookList := []types.Book{}

	rows, err := db.Query("SELECT * FROM books WHERE Status = 'requested'")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Title, &book.Author, &book.BookStatus, &book.Quantity)
		if err != nil {
			fmt.Println(err)
		}

		BookList = append(BookList, book)
	}

	fmt.Println(BookList)
}

func ApproveBookRequestPost(requestID int, bookID int) error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	_, err = db.Query("UPDATE requests SET book_status = 'approved' WHERE request_id = ?", requestID)
	if err != nil {
		return fmt.Errorf("error updating book: %s", err)
	}

	_, err = db.Query("UPDATE books SET status = 'not available' WHERE book_id = ?", bookID)
	if err != nil {
		return fmt.Errorf("error updating book: %s", err)
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
		err := rows.Scan(&request.RequestID, &request.BookID, &request.UserID, &request.RequestStatus)
		if err != nil {
			return nil, fmt.Errorf("error scanning request rows: %s", err)
		}
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

	// fmt.Println("checking")

	var RequestList []types.RequestAdminView

	rows, err := db.Query("SELECT * FROM requests WHERE book_status = 'rejected'")
	if err != nil {
		return nil, fmt.Errorf("error searching requests: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var request types.RequestAdminView
		err := rows.Scan(&request.RequestID, &request.BookID, &request.UserID, &request.RequestStatus)
		if err != nil {
			return nil, fmt.Errorf("error scanning request rows: %s", err)
		}
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

	// fmt.Println(RequestList)
	return RequestList, nil
}
