package models

import (
	"fmt"
	"mvc/pkg/types"
)

func GetBookTitleByBookID(bookId int) (string, error) {
	db, err := Connect()
	if err != nil {
		return "", fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var title string
	err = db.QueryRow("SELECT title FROM books WHERE book_id = ?", bookId).Scan(&title)
	if err != nil {
		return "", fmt.Errorf("error searching book: %s", err)
	}

	return title, nil
}

func GetUserIDFromUsername(username string) (int, error) {
	db, err := Connect()
	if err != nil {
		return 0, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var userId int
	err = db.QueryRow("SELECT user_id FROM users WHERE username = ?", username).Scan(&userId)
	if err != nil {
		return 0, fmt.Errorf("error searching user: %s", err)
	}

	return userId, nil
}

func GetUsernameFromID(userId int) (string, error) {
	db, err := Connect()
	if err != nil {
		return "", fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var username string
	err = db.QueryRow("SELECT username FROM users WHERE user_id = ?", userId).Scan(&username)
	if err != nil {
		return "", fmt.Errorf("error searching user: %s", err)
	}

	return username, nil
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

func GetBookIdByRequestId(requestId int) (int, error) {
	db, err := Connect()
	if err != nil {
		return 0, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var bookId int
	err = db.QueryRow("SELECT book_id FROM requests WHERE request_id = ?", requestId).Scan(&bookId)
	if err != nil {
		return 0, fmt.Errorf("error searching book: %s", err)
	}

	return bookId, nil
}

func GetUserIdByBookId(bookId int) (int, error) {
	db, err := Connect()
	if err != nil {
		return 0, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	var userId int
	err = db.QueryRow("SELECT user_id FROM requests WHERE book_id = ?", bookId).Scan(&userId)
	if err != nil {
		return 0, fmt.Errorf("error searching user: %s", err)
	}

	return userId, nil
}
