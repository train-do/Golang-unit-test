package service

import (
	"be-golang-chapter-36-implem/model"

	"github.com/stretchr/testify/mock"
)

type CustomerServiceMock struct {
	mock.Mock
}

func (customerServiceMock *CustomerServiceMock) Create(customer *model.Customer) error {
	return nil
}

func (customerServiceMock *CustomerServiceMock) GetAll() (*[]model.Customer, error) {
	return nil, nil
}

func (customerServiceMock *CustomerServiceMock) Login(customer model.Customer) (*model.Customer, error) {
	args := customerServiceMock.Called(customer)
	if customerResult := args.Get(0); customerResult != nil {
		return customerResult.(*model.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}
