package utils

import (
	"encoding/json"
	"net/http"
)

type Pagination struct {
	CurrentPage  int `json:"current_page"`
	Limit        int `json:"limit"`
	TotalPages   int `json:"total_pages"`
	TotalRecords int `json:"total_records"`
}

type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type ResponseValidationError struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func WriteSuccess(w http.ResponseWriter, message string, statusCode int, data interface{}, pagination *Pagination) {
	writeJSON(w, statusCode, Response{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	})
}

func WriteError(w http.ResponseWriter, message string, statusCode int) {
	writeJSON(w, statusCode, Response{
		Success: false,
		Message: message,
		Data:    nil,
	})
}

func ResponseErrorValidation(w http.ResponseWriter, statusCode int, message string, error interface{}) {
	response := ResponseValidationError{
		Success: false,
		Message: message,
		Errors:  error,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func writeJSON(w http.ResponseWriter, statusCode int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}
