package models

func DeleteBookPost(bookID int) error {
	db, err := Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Query("DELETE FROM requests WHERE book_id = ?", bookID)
	if err != nil {
		return err
	}

	_, err = db.Query("DELETE FROM books WHERE book_id = ?", bookID)
	if err != nil {
		return err
	}

	return nil
}
