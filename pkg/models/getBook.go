package models

import (
	"fmt"
	"mvc/pkg/types"
)

func GetBook() error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	BookList := []types.Book{}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return fmt.Errorf("error querying DB: %s", err)
	}

	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Title, &book.Author, &book.BookStatus, &book.Quantity)
		if err != nil {
			return fmt.Errorf("error scanning rows: %s", err)
		}

		BookList = append(BookList, book)
	}

	fmt.Println(BookList)
	// return BookList
	return nil

}
