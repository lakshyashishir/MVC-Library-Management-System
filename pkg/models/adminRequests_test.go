package models

import (
	"mvc/pkg/types"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestViewAdmins(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"user_id", "username", "hash", "salt", "role"}).
		AddRow(1, "check", "hash", "salt", "admin requested").
		AddRow(2, "Lakshya", "hash", "salt", "admin")

	mock.ExpectQuery("SELECT user_id,username, hash, salt, role FROM users WHERE role = 'admin requested'").
		WillReturnRows(rows)

	adminRequests, err := GetAdminRequests(db)
	if err != nil {
		t.Errorf("Error while getting admin requests: %s", err)
		return
	}

	expected := []types.User{
		{
			UserID:   1,
			Username: "check",
			Hash:     "hash",
			Salt:     "salt",
			Role:     "admin requested",
		},
		{
			UserID:   2,
			Username: "Lakshya",
			Hash:     "hash",
			Salt:     "salt",
			Role:     "admin",
		},
	}

	assert.Equal(t, expected, adminRequests)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
