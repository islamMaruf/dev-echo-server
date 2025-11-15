package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

// HTTPError represents a structured HTTP error response.
type HTTPError struct {
	Status  int    `json:"status"`  // HTTP status code
	Message string `json:"message"` // Error message
	Expose  bool   `json:"-"`       // Whether to expose details to client
}

func (e *HTTPError) Error() string {
	return e.Message
}

// NewHTTPError creates a new HTTPError with the specified status, message, and expose flag.
//
// Parameters:
//   - status: HTTP status code (e.g., 404, 500)
//   - message: Human-readable error message
//   - expose: Whether the error should be exposed to the client
//
// Returns:
//   - *HTTPError: A new error instance
func NewHTTPError(status int, message string, expose bool) *HTTPError {
	return &HTTPError{
		Status:  status,
		Message: message,
		Expose:  expose,
	}
}

// NotFound creates a 404 Not Found error.
func NotFound() *HTTPError {
	return NewHTTPError(http.StatusNotFound, "Not Found", true)
}

// InternalServerError creates a 500 Internal Server Error.
func InternalServerError() *HTTPError {
	return NewHTTPError(http.StatusInternalServerError, "Internal Server Error", false)
}

// ErrorHandler returns a middleware that handles errors and panics gracefully.
// It recovers from panics, logs the error, and returns a JSON error response.
//
// Parameters:
//   - next: The next handler in the middleware chain
//
// Returns:
//   - http.Handler: Handler with error handling and panic recovery
//
// On panic, it returns a 500 Internal Server Error with JSON response.
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
