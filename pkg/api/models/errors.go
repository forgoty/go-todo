package models

import "fmt"

type APIError struct {
	Message string `json:"message,omitempty"`
}

func (e *APIError) ToString() string {
	return fmt.Sprintf("Error: %s", e.Message)
}
