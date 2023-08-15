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
		err := rows.Scan(&book.BookID, &book.Title, &book.Author, &book.BookStatus, &book.Quantity)
		if err != nil {
			return nil, fmt.Errorf("error scanning book rows: %s", err)
		}

		BookList = append(BookList, book)
	}

	return BookList, nil
}

func UserReturnBookPost(BookID int, UserID int) error {
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer db.Close()

	rows, err := db.Query("Select * from books where book_id = ? AND status = 'available'", BookID)
	if err != nil {
		err = fmt.Errorf("error checking book status: %s", err)
		return err
	}
	if rows == nil {
		_, err = db.Query("UPDATE books SET status = 'available' WHERE book_id = ?", BookID)
		if err != nil {
			return fmt.Errorf("error updating book status: %s", err)
		}
	}

	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookID, &book.Title, &book.Author, &book.BookStatus, &book.Quantity)
		if err != nil {
			return fmt.Errorf("error scanning book rows: %s", err)
		}

		_, err = db.Query("UPDATE books SET quantity = quantity + 1 WHERE book_id = ?", BookID)
		if err != nil {
			return fmt.Errorf("error updating book quantity: %s", err)
		}
	}

	_, err = db.Exec("DELETE from requests where book_id = ? AND user_id = ?", BookID, UserID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func UserRequestBookPost(UserID int, BookID int) error {
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO requests (user_id, book_id, book_status) VALUES (?, ?, 'pending')", UserID, BookID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
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

func GetUserBooks(UserID int) ([]types.BookUserView, error) {
	db, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	BookList := []types.BookUserView{}

	rows, err := db.Query("Select * from requests where book_status = 'approved' and user_id = ?", UserID)
	if err != nil {
		return nil, fmt.Errorf("error searching books: %s", err)
	}

	for rows.Next() {
		var book types.BookUserView
		err := rows.Scan(&book.RequestID, &book.UserID, &book.BookID, &book.BookStatus)
		if err != nil {
			return nil, fmt.Errorf("error scanning book rows: %s", err)
		}
		book.Title, err = GetBookTitleByBookID(book.BookID)
		if err != nil {
			return nil, fmt.Errorf("error getting book title: %s", err)
		}
		BookList = append(BookList, book)
	}

	return BookList, nil
}

func GetUserRequestsPending(UserID int) ([]types.BookUserView, error) {
	db, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	BookList := []types.BookUserView{}

	rows, err := db.Query("Select * from requests where book_status = 'pending' and user_id = ?", UserID)
	if err != nil {
		return nil, fmt.Errorf("error searching books: %s", err)
	}

	for rows.Next() {
		var book types.BookUserView
		err := rows.Scan(&book.RequestID, &book.UserID, &book.BookID, &book.BookStatus)
		if err != nil {
			return nil, fmt.Errorf("error scanning book rows: %s", err)
		}
		book.Title, err = GetBookTitleByBookID(book.BookID)
		if err != nil {
			return nil, fmt.Errorf("error getting book title: %s", err)
		}
		BookList = append(BookList, book)
	}

	return BookList, nil
}

func GetUserRequestsRejected(UserID int) ([]types.BookUserView, error) {
	db, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	BookList := []types.BookUserView{}

	rows, err := db.Query("Select * from requests where book_status = 'rejected' and user_id = ?", UserID)
	if err != nil {
		return nil, fmt.Errorf("error searching books: %s", err)
	}

	for rows.Next() {
		var book types.BookUserView
		err := rows.Scan(&book.RequestID, &book.UserID, &book.BookID, &book.BookStatus)
		if err != nil {
			return nil, fmt.Errorf("error scanning book rows: %s", err)
		}
		book.Title, err = GetBookTitleByBookID(book.BookID)
		if err != nil {
			return nil, fmt.Errorf("error getting book title: %s", err)
		}
		BookList = append(BookList, book)
	}

	return BookList, nil
}
