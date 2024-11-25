package helper

import (
	"encoding/json"
	"net/http"
)

type Respose struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func SuccessResponse(w http.ResponseWriter, message string, statusCode int) {
	respon := Respose{
		Status:  true,
		Message: message,
	}
	w.WriteHeader(statusCode)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(respon)
}

func BadResponse(w http.ResponseWriter, message string, statusCode int) {
	respon := Respose{
		Status:  false,
		Message: message,
	}
	w.WriteHeader(statusCode)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(respon)
}

func SuccessResponseWithData(w http.ResponseWriter, message string, statusCode int, data interface{}) {
	respon := Respose{
		Status:  true,
		Message: message,
		Data:    data,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(respon)
}
