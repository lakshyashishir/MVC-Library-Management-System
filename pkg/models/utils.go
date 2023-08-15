package models

import (
	"fmt"
	"mvc/pkg/types"
)

func GetBookTitleByBookID(bookID int) (string, error) {
	db, err := Connect()
	if err != nil {
		return "", fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var Title string
	err = db.QueryRow("SELECT title FROM books WHERE book_id = ?", bookID).Scan(&Title)
	if err != nil {
		return "", fmt.Errorf("error searching book: %s", err)
	}

	return Title, nil
}

func GetUserIDFromUsername(username string) (int, error) {
	db, err := Connect()
	if err != nil {
		return 0, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var UserID int
	err = db.QueryRow("SELECT user_id FROM users WHERE username = ?", username).Scan(&UserID)
	if err != nil {
		return 0, fmt.Errorf("error searching user: %s", err)
	}

	return UserID, nil
}

func GetUsernameFromID(userID int) (string, error) {
	db, err := Connect()
	if err != nil {
		return "", fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var Username string
	err = db.QueryRow("SELECT username FROM users WHERE user_id = ?", userID).Scan(&Username)
	if err != nil {
		return "", fmt.Errorf("error searching user: %s", err)
	}

	return Username, nil
}

func GetBookStatusByBookID(bookID int) (types.BookStatus, error) {
	db, err := Connect()
	if err != nil {
		return "", fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var book types.Book

	err = db.QueryRow("SELECT * FROM books WHERE book_id = ?", bookID).Scan(&book.BookID, &book.Title, &book.Author, &book.BookStatus, &book.Quantity)
	if err != nil {
		return "", fmt.Errorf("error scanning book rows: %s", err)
	}

	return book.BookStatus, nil
}

func GetBookIdByRequestId(RequestID int) (int, error) {
	db, err := Connect()
	if err != nil {
		return 0, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var BookID int
	err = db.QueryRow("SELECT book_id FROM requests WHERE request_id = ?", RequestID).Scan(&BookID)
	if err != nil {
		return 0, fmt.Errorf("error searching book: %s", err)
	}

	return BookID, nil
}

func GetUserIdByBookId(BookID int) (int, error) {
	db, err := Connect()
	if err != nil {
		return 0, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var UserID int
	err = db.QueryRow("SELECT user_id FROM requests WHERE book_id = ?", BookID).Scan(&UserID)
	if err != nil {
		return 0, fmt.Errorf("error searching user: %s", err)
	}

	return UserID, nil
}
