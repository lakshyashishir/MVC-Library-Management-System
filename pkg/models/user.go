package models

import (
	"fmt"
	"mvc/pkg/types"
)

func GetBook() ([]types.Book, error) {
	db, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	BookList := []types.Book{}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, fmt.Errorf("error searching books: %s", err)
	}

	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Title, &book.Author, &book.BookStatus, &book.Quantity)
		if err != nil {
			return nil, fmt.Errorf("error scanning book rows: %s", err)
		}

		BookList = append(BookList, book)
	}

	return BookList, nil
}

func UserReturnBookPost(BookID int, UserID int) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	_, err = db.Exec("UPDATE books SET status = 'available' WHERE book_id = ?", BookID)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec("DELETE from requests where book_id = ? AND user_id = ?", BookID, UserID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UserRequestBookPost(BookID int, UserID int) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	_, err = db.Exec("UPDATE books SET status = 'requested' WHERE book_id = ?", BookID)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec("INSERT INTO requests (user_id, book_id, book_status) VALUES (?, ?, 'pending')", UserID, BookID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UserRemoveRequestPost(RequestID int, BookID int) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	_, err = db.Exec("UPDATE books SET status = 'available' WHERE book_id = ?", BookID)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec("DELETE from requests where request_id = ?", RequestID)
	if err != nil {
		fmt.Println(err)
		return
	}
}
