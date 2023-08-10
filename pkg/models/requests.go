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

func ApproveBookPost(request_id int) error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	_, err = db.Query("UPDATE requests SET book_status = 'approved' WHERE request_id = ?", request_id)
	if err != nil {
		return fmt.Errorf("error updating book: %s", err)
	}

	return nil
}

func RejectBookPost(title string) error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	_, err = db.Query("DELETE FROM books WHERE Title = ?", title)
	if err != nil {
		return fmt.Errorf("error deleting book: %s", err)
	}

	return nil
}
