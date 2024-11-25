package handler

import (
	"be-golang-chapter-36-implem/helper"
	"be-golang-chapter-36-implem/model"
	"be-golang-chapter-36-implem/service"

	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type CustomerHadler struct {
	Service service.AllService
	Log     *zap.Logger
}

func NewCustomerHandler(service service.AllService, log *zap.Logger) CustomerHadler {
	return CustomerHadler{
		Service: service,
		Log:     log,
	}
}

func (CustomerHadler *CustomerHadler) Create(w http.ResponseWriter, r *http.Request) {

}

func (CustomerHadler *CustomerHadler) GetAll(w http.ResponseWriter, r *http.Request) {
	customers, err := CustomerHadler.Service.CustomerService.GetAll()
	if err != nil {
		CustomerHadler.Log.Error("error get all data :", zap.Error(err))
		helper.BadResponse(w, err.Error(), http.StatusBadRequest)
	}

	helper.SuccessResponseWithData(w, "success", http.StatusOK, customers)
}

func (CustomerHadler *CustomerHadler) Login(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var customer model.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		CustomerHadler.Log.Error("Failed to parse login request", zap.Error(err))
		helper.BadResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if customer.Email == "" || customer.Password == "" {
		CustomerHadler.Log.Warn("Missing email or password")
		helper.BadResponse(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Call service to authenticate user
	_, err := CustomerHadler.Service.CustomerService.Login(customer)
	if err != nil {
		CustomerHadler.Log.Error("Authentication failed", zap.Error(err))
		if err.Error() == "invalid email" {
			helper.BadResponse(w, "Invalid email", http.StatusUnauthorized)
		} else {
			helper.BadResponse(w, "Invalid password", http.StatusUnauthorized)
		}
		return
	}

	// Respond with the token
	response := map[string]string{
		"token": "123456",
	}

	helper.SuccessResponseWithData(w, "Success Login", http.StatusOK, response)
}
