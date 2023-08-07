package models

import (
	"fmt"
	"mvc/pkg/types"
)

func AddBook(book types.Book) error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	book.BookStatus = "available"

	_, err = db.Exec("INSERT INTO books (Title, Author, BookStatus, Quantity) VALUES (?, ?, ?, ?)", book.Title, book.Author, book.BookStatus, book.Quantity)
	if err != nil {
		return fmt.Errorf("error inserting into DB: %s", err)
	}

	return nil
}
