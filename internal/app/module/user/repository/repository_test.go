package repository_test

import (
	"CodeWithAzri/internal/app/module/user/entity"
	"CodeWithAzri/internal/app/module/user/repository"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func initializeMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *repository.Repository) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}

	mockRepo := repository.NewRepository(db)

	return db, mock, mockRepo
}

func TestRepository_Create(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	user := &entity.User{
		ID:        "1",
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: 121212,
		UpdatedAt: 121212,
	}

	mock.ExpectExec("INSERT INTO users (id, name, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)").
		WithArgs(user.ID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Create(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestRepository_ReadMany(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow("1", "John Doe", "john@domain.com", 121212, 121212).
		AddRow("2", "Jane Doe", "jane@domain.com", 121212, 121212)

	mock.ExpectQuery("SELECT * FROM users LIMIT $1 OFFSET $2").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(rows)

	users, err := repo.ReadMany(10, 0)
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_ReadOne(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	userID := "1"
	user := &entity.User{
		ID:        userID,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: 121212,
		UpdatedAt: 121212,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow(user.ID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt)

	mock.ExpectQuery(`SELECT * FROM users WHERE id = $1`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(rows)

	result, err := repo.ReadOne(userID)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_Update(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	userID := "1"
	user := &entity.User{
		ID:        userID,
		Name:      "Updated Name",
		Email:     "john@example.com",
		CreatedAt: 121212,
		UpdatedAt: 121212,
	}

	mock.ExpectExec("UPDATE users SET name = $1 WHERE id = $2").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Update(userID, user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_Delete(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	userID := "1"

	mock.ExpectExec("DELETE FROM users WHERE id = $1").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(userID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_ReadOne_Error(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	userID := "1"

	mock.ExpectQuery(`SELECT * FROM users WHERE id = $1`).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(userID, "John Doe").
			RowError(0, sql.ErrTxDone))

	_, err := repo.ReadOne(userID)
	assert.Error(t, err)
	assert.ErrorIs(t, err, sql.ErrTxDone)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_ReadOne_NoRecord(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	userID := "nonexistent"

	mock.ExpectQuery(`SELECT * FROM users WHERE id = $1`).
		WithArgs(userID).
		WillReturnError(sql.ErrNoRows)

	result, err := repo.ReadOne(userID)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, sql.ErrNoRows)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_Delete_NoRecord(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	userID := "nonexistent"

	mock.ExpectExec("DELETE FROM users WHERE id = $1").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 0)).
		WillReturnError(sql.ErrNoRows)

	err := repo.Delete(userID)
	assert.Error(t, err)
	assert.ErrorIs(t, err, sql.ErrNoRows)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_ReadMany_Error(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	mock.ExpectQuery("SELECT * FROM users LIMIT $1 OFFSET $2").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(sql.ErrTxDone)

	_, err := repo.ReadMany(10, 0)
	assert.Error(t, err)
	assert.ErrorIs(t, err, sql.ErrTxDone)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_ReadMany_ErrorScan(t *testing.T) {
	db, mock, repo := initializeMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow("1", "John Doe", "john@domain.com", "invalid_created_at", 121212).
		AddRow("2", "Jane Doe", "jane@domain.com", 121212, 121212)

	mock.ExpectQuery("SELECT * FROM users LIMIT $1 OFFSET $2").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(rows)

	users, err := repo.ReadMany(10, 0)
	assert.Error(t, err)
	assert.Nil(t, users)
	assert.Contains(t, err.Error(), "convert")
	assert.NoError(t, mock.ExpectationsWereMet())
}
