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
