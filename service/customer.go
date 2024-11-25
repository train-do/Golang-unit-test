package service

import (
	"be-golang-chapter-36-implem/helper"
	"be-golang-chapter-36-implem/model"
	"be-golang-chapter-36-implem/repository"
	"errors"

	"go.uber.org/zap"
)

type CustomerServiceInterface interface {
	Create(customer *model.Customer) error
	GetAll() (*[]model.Customer, error)
	Login(customer model.Customer) (*model.Customer, error)
}

type CustomerService struct {
	Repo repository.AllRepository
	Log  *zap.Logger
}

func NewCustomerService(repo repository.AllRepository, log *zap.Logger) CustomerServiceInterface {
	return &CustomerService{
		Repo: repo,
		Log:  log,
	}
}

func (customerService *CustomerService) Create(customer *model.Customer) error {
	hashedPassword, err := helper.HashPassword(customer.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}
	customer.Password = hashedPassword
	return customerService.Repo.CustomerRep.Create(customer)
}

func (customerService *CustomerService) GetAll() (*[]model.Customer, error) {
	return customerService.Repo.CustomerRep.GetAll()
}

func (customerService *CustomerService) Login(customer model.Customer) (*model.Customer, error) {
	// check by email
	customerResult, err := customerService.Repo.CustomerRep.GetByCondition(customer)
	if err != nil {
		return nil, errors.New("invalid email")
	}

	// Verify password (assuming password is hashed and needs comparison)
	if !helper.CheckPassword(customer.Password, customerResult.Password) {
		return nil, errors.New("invalid password")
	}

	return customerResult, nil
}
