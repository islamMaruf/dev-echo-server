// Package middleware provides HTTP middleware for logging, security, and error handling.
package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logger returns a middleware that logs HTTP requests to a file and console.
// It captures request details including method, URI, status code, request body,
// and response time. Each request is assigned a unique request ID for tracking.
//
// Parameters:
//   - logFile: File to write JSON-formatted logs to. Pass nil to skip file logging.
//
// Returns:
//   - A middleware function that wraps an http.Handler
//
// Logs include:
//   - timestamp: ISO 8601 formatted timestamp
//   - method: HTTP method (GET, POST, etc.)
//   - URI: Request URI
//   - status: HTTP response status code
//   - requestBody: Parsed JSON request body (if present)
//   - responseTime: Request duration in milliseconds
//
// Console logging is enabled only when NODE_ENV=development
func Logger(logFile *os.File) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Generate request ID
			requestID := uuid.New().String()
			r.Header.Set("X-Request-ID", requestID)

			// Wrap response writer to capture status code
			wrapped := &responseWriter{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			// Read body for logging
			var bodyData interface{}
			if r.Body != nil {
				bodyBytes, err := io.ReadAll(r.Body)
				if err == nil && len(bodyBytes) > 0 {
					json.Unmarshal(bodyBytes, &bodyData)
					// Restore body for handlers
					r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
				}
			}

			next.ServeHTTP(wrapped, r)

			// Log to file
			logEntry := map[string]interface{}{
				"timestamp":    time.Now().Format(time.RFC3339),
				"method":       r.Method,
				"URI":          r.URL.String(),
				"status":       wrapped.status,
				"requestBody":  bodyData,
				"responseTime": time.Since(start).Milliseconds(),
			}

			logJSON, _ := json.Marshal(logEntry)
			if logFile != nil {
				logFile.Write(append(logJSON, '\n'))
			}

			// Also log to console in dev mode
			if os.Getenv("NODE_ENV") == "development" {
				log.Printf("%s %s %d %dms\n", r.Method, r.URL.Path, wrapped.status, time.Since(start).Milliseconds())
			}
		})
	}
}
