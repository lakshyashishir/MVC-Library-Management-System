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
