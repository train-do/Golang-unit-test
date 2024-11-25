package repository

import (
	"be-golang-chapter-36-implem/model"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

type CustomerRepoInterface interface {
	Create(customer *model.Customer) error
	GetAll() (*[]model.Customer, error)
	GetByCondition(customer model.Customer) (*model.Customer, error)
}

type CustomerRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewCustomerRepository(db *sql.DB, log *zap.Logger) CustomerRepoInterface {
	return &CustomerRepository{
		DB:     db,
		Logger: log,
	}
}

func (customerRepo *CustomerRepository) Create(customer *model.Customer) error {
	query := "INSERT INTO customers (name, email, phone, password) VALUES ($1, $2, $3, $4) RETURNING id"
	err := customerRepo.DB.QueryRow(query, customer.Name, customer.Email, customer.Phone, customer.Password).Scan(&customer.ID)
	return err
}

func (customerRepo *CustomerRepository) GetAll() (*[]model.Customer, error) {
	var customers []model.Customer
	query := "SELECT name, email, phone, password FROM customers"
	rows, err := customerRepo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.Customer
		if err := rows.Scan(&user.Name, &user.Email, &user.Phone, &user.Password); err != nil {
			return nil, err
		}
		customers = append(customers, user)
	}
	return &customers, nil
}

func (CustomerRepository *CustomerRepository) GetByCondition(customer model.Customer) (*model.Customer, error) {
	var customerResult model.Customer

	query := "SELECT id, name, email, phone, password FROM customers WHERE 1=1"
	args := []interface{}{}

	if customer.Email != "" {
		query += " AND email = ?"
		args = append(args, customer.Email)
	}

	if customer.Phone != "" {
		query += " AND phone = ?"
		args = append(args, customer.Phone)
	}

	row := CustomerRepository.DB.QueryRow(query, args...)
	err := row.Scan(&customerResult.ID, &customerResult.Name, &customerResult.Email, &customerResult.Phone, &customerResult.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Return nil if no rows found
		}
		return nil, err
	}

	return &customerResult, nil
}
