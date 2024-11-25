package handler

import (
	"be-golang-chapter-36-implem/helper"
	"be-golang-chapter-36-implem/model"
	"be-golang-chapter-36-implem/service"
	"bytes"
	"errors"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCustomerHandler_Login(t *testing.T) {

	mockService := service.CustomerServiceMock{}

	allService := service.AllService{
		CustomerService: &mockService,
	}

	logger := zap.NewNop() // No-op logger for testing
	customerHadler := NewCustomerHandler(allService, logger)

	tests := []struct {
		name           string
		requestBody    interface{}
		arg1MockSetup  *model.Customer
		arg2MockSetup  error
		expectedStatus int
		expectedBody   helper.Respose
	}{
		{
			name: "Success Login",
			requestBody: model.Customer{
				Email:    "test@example.com",
				Password: "password123",
			},
			arg1MockSetup: &model.Customer{
				Email:    "test@example.com",
				Password: "password123",
			},
			arg2MockSetup:  nil,
			expectedStatus: http.StatusOK,
			expectedBody: helper.Respose{
				Status:  true,
				Message: "Success Login",
			},
		},
		{
			name:           "Invalid Request Body",
			requestBody:    "invalid-json",
			arg1MockSetup:  nil,
			arg2MockSetup:  nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody: helper.Respose{
				Status:  false,
				Message: "Invalid request body",
			},
		},
		{
			name: "Missing Email or Password",
			requestBody: model.Customer{
				Email:    "",
				Password: "",
			},
			arg1MockSetup:  nil,
			arg2MockSetup:  nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody: helper.Respose{
				Status:  false,
				Message: "Email and password are required",
			},
		},
		{
			name: "Authentication Failed Email",
			requestBody: model.Customer{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
			arg1MockSetup:  nil,
			arg2MockSetup:  errors.New("invalid email"),
			expectedStatus: http.StatusUnauthorized,
			expectedBody: helper.Respose{
				Status:  false,
				Message: "Invalid email",
			},
		},
		{
			name: "Authentication Failed Password",
			requestBody: model.Customer{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
			arg1MockSetup:  nil,
			arg2MockSetup:  errors.New("invalid password"),
			expectedStatus: http.StatusUnauthorized,
			expectedBody: helper.Respose{
				Status:  false,
				Message: "Invalid password",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare request
			requestBody, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err, "Failed to marshal request body")

			// request
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))

			// Prepare response recorder
			w := httptest.NewRecorder()

			// setup mock
			mockService.On("Login", tt.requestBody).Once().Return(tt.arg1MockSetup, tt.arg2MockSetup)

			// Call the handler
			customerHadler.Login(w, req)

			// Assert status code
			res := w.Result()
			assert.Equal(t, tt.expectedStatus, res.StatusCode)

			// Periksa body response
			var responseBody helper.Respose
			err = json.NewDecoder(w.Body).Decode(&responseBody)
			assert.NoError(t, err, "Failed to decode response body")
			assert.Equal(t, tt.expectedBody.Status, responseBody.Status)
			assert.Equal(t, tt.expectedBody.Message, responseBody.Message)
		})
	}
}
