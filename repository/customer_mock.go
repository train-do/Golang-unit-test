package repository

import (
	"be-golang-chapter-36-implem/model"

	"github.com/stretchr/testify/mock"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (customerRepositoryMock *CustomerRepositoryMock) Create(customer *model.Customer) error {
	return nil
}

func (customerRepositoryMock *CustomerRepositoryMock) GetAll() (*[]model.Customer, error) {
	return nil, nil
}

func (customerRepositoryMock *CustomerRepositoryMock) GetByCondition(customer model.Customer) (*model.Customer, error) {
	args := customerRepositoryMock.Called(customer)
	if customerResult := args.Get(0); customerResult != nil {
		return customerResult.(*model.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}
