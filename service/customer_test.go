package service

import (
	"be-golang-chapter-36-implem/helper"
	"be-golang-chapter-36-implem/model"
	"be-golang-chapter-36-implem/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// logger := zap.NewNop()

	mockRepo := repository.CustomerRepositoryMock{}
	allRepo := repository.AllRepository{
		CustomerRep: &mockRepo,
	}

	customerService := NewCustomerService(allRepo, nil)

	// Mock data
	validCustomer := model.Customer{
		Email:    "test@example.com",
		Password: "validpassword",
	}

	hashedPassword, _ := helper.HashPassword("validpassword")
	storedCustomer := &model.Customer{
		Email:    "test@example.com",
		Password: hashedPassword,
	}

	hashedPasswordFailed, _ := helper.HashPassword("1231313")
	storedCustomerFailed := &model.Customer{
		Email:    "test@example.com",
		Password: hashedPasswordFailed,
	}

	t.Run("invalid email", func(t *testing.T) {
		// Setup mock
		// mockRepo.On("GetByCondition", validCustomer).Return(nil, errors.New("not found"))
		mockRepo.On("GetByCondition", validCustomer).Once().Return(nil, errors.New("not found"))
		// Test
		result, err := customerService.Login(validCustomer)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "invalid email", err.Error())

		mockRepo.AssertExpectations(t)
	})

	t.Run("successful login", func(t *testing.T) {
		// mockRepo.ExpectedCalls = nil // untuk reset mock setup sebelumnya
		// Setup mock
		// mockRepo.On("GetByCondition", validCustomer).Return(storedCustomer, nil)
		mockRepo.On("GetByCondition", validCustomer).Once().Return(storedCustomer, nil) // untuk digunakan sekali

		// Test
		result, err := customerService.Login(validCustomer)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, validCustomer.Email, result.Email)

		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid password", func(t *testing.T) {
		// mockRepo.ExpectedCalls = nil // untuk reset mock setup sebelumnya
		// Setup mock
		// mockRepo.On("GetByCondition", validCustomer).Return(storedCustomerFailed, nil)
		mockRepo.On("GetByCondition", validCustomer).Once().Return(storedCustomerFailed, nil) // untuk digunakan sekali

		// Test
		result, err := customerService.Login(validCustomer)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "invalid password", err.Error())

		mockRepo.AssertExpectations(t)
	})
}
