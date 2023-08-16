package models

func DeleteBookPost(bookId int) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Query("DELETE FROM requests WHERE book_id = ?", bookId)
	if err != nil {
		return err
	}

	_, err = db.Query("DELETE FROM books WHERE book_id = ?", bookId)
	if err != nil {
		return err
	}

	return nil
}
