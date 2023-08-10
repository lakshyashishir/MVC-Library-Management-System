package models

import (
	"crypto/rand"
	"fmt"
)

func GeneratSessionID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("error generating sessionID: %s", err)
	}
	return fmt.Sprintf("%x", b), nil
}

func UpdateSessionID(UserID int, sessionID string) error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	_, err = db.Query("DELETE FROM cookies WHERE userId = ?", UserID)
	if err != nil {
		return fmt.Errorf("error updating sessionID: %s", err)
	}

	_, err = db.Query("INSERT INTO cookies (sessionId, userId) VALUES (?, ?)", sessionID, UserID)
	if err != nil {
		return fmt.Errorf("error updating sessionID: %s", err)
	}

	return nil
}

func DeleteSessionID(session_id string) error {
	db, err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	defer db.Close()

	_, err = db.Query("DELETE FROM cookies WHERE sessionId = ?", session_id)
	if err != nil {
		return fmt.Errorf("error deleting session: %s", err)
	}

	return nil
}
