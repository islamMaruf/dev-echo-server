package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

type HTTPError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Expose  bool   `json:"-"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

// NewHTTPError creates a new HTTP error
func NewHTTPError(status int, message string, expose bool) *HTTPError {
	return &HTTPError{
		Status:  status,
		Message: message,
		Expose:  expose,
	}
}

// NotFound creates a 404 error
func NotFound() *HTTPError {
	return NewHTTPError(http.StatusNotFound, "Not Found", true)
}

// InternalServerError creates a 500 error
func InternalServerError() *HTTPError {
	return NewHTTPError(http.StatusInternalServerError, "Internal Server Error", false)
}

// ErrorHandler middleware handles errors and panics
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(InternalServerError())
			}
		}()

		next.ServeHTTP(w, r)
	})
}
