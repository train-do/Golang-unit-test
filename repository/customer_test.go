package repository

import (
	"be-golang-chapter-36-implem/model"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	return db, mock
}

func TestCustomerRepository_GetByCondition_GORM(t *testing.T) {
	db, mock := setupTestDB(t)
	defer func() { _ = mock.ExpectationsWereMet() }()

	customerRepo := NewCustomerRepository(db, nil)

	t.Run("successfully get customer by email", func(t *testing.T) {
		customer := model.Customer{
			Email: "johndoe@example.com",
		}

		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "password"}).
			AddRow(1, "John Doe", "johndoe@example.com", "123456789", "password123")

		mock.ExpectQuery(`SELECT id, name, email, phone, password FROM customers WHERE 1=1 AND email = ?`).
			WithArgs(customer.Email). // GORM secara default mengisi LIMIT dengan 1
			WillReturnRows(rows)

		result, err := customerRepo.GetByCondition(customer)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "John Doe", result.Name)
		assert.Equal(t, "johndoe@example.com", result.Email)
	})

	t.Run("successfully get customer by phone", func(t *testing.T) {
		customer := model.Customer{
			Phone: "123456789",
		}

		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "password"}).
			AddRow(1, "John Doe", "johndoe@example.com", "123456789", "password123")

		mock.ExpectQuery(`SELECT id, name, email, phone, password FROM customers WHERE 1=1 AND phone = ?`).
			WithArgs(customer.Phone).
			WillReturnRows(rows)

		result, err := customerRepo.GetByCondition(customer)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "John Doe", result.Name)
		assert.Equal(t, "123456789", result.Phone)
	})

	t.Run("customer not found", func(t *testing.T) {
		customer := model.Customer{
			Email: "unknown@example.com",
		}

		mock.ExpectQuery(`SELECT id, name, email, phone, password FROM customers WHERE 1=1 AND email = ?`).
			WithArgs(customer.Email).
			WillReturnError(gorm.ErrRecordNotFound)

		result, err := customerRepo.GetByCondition(customer)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}

func BenchmarkGetAll(b *testing.B) {
	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		b.Fatalf("Failed to create mock DB: %s", err)
	}
	defer db.Close()

	// Mock data
	mockRows := sqlmock.NewRows([]string{"name", "email", "phone", "password"}).
		AddRow("John Doe", "john@example.com", "123456789", "hashedPassword").
		AddRow("Jane Doe", "jane@example.com", "987654321", "hashedPassword")

	// Tambahkan ekspektasi untuk setiap iterasi benchmark
	mock.MatchExpectationsInOrder(false) // Tidak peduli urutan ekspektasi

	// Repository setup
	logger := zap.NewNop()
	repo := NewCustomerRepository(db, logger)

	for i := 0; i < b.N; i++ {
		mock.ExpectQuery("SELECT name, email, phone, password FROM customers").
			WillReturnRows(mockRows)

		customers, err := repo.GetAll()
		if err != nil {
			b.Errorf("Unexpected error: %v", err)
		}

		if customers == nil {
			b.Errorf("Expected value not to be nil")
		}
	}
}
