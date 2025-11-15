// Package routes provides HTTP route handlers for the dev echo server.
// It includes a welcome endpoint and a catch-all echo handler that mirrors request data.
package routes

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterRoutes registers all application routes with the provided HTTP multiplexer.
// It configures:
//   - GET / - Returns a welcome message
//   - * /* - Catch-all route that echoes the request body back to the client
//
// Parameters:
//   - mux: The HTTP multiplexer to register routes with
//
// Example:
//
//	mux := http.NewServeMux()
//	routes.RegisterRoutes(mux)
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", combinedHandler)
}

func combinedHandler(w http.ResponseWriter, r *http.Request) {
	// Home route - exact match for "/"
	if r.URL.Path == "/" {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		response := map[string]string{
			"message": "Welcome",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Catch-all for other paths
	catchAllHandler(w, r)
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {

	var bodyData interface{}
	if r.Body != nil {
		bodyBytes, err := io.ReadAll(r.Body)
		if err == nil && len(bodyBytes) > 0 {
			json.Unmarshal(bodyBytes, &bodyData)
		}
	}

	response := map[string]interface{}{
		"response": map[string]interface{}{
			"data":    bodyData,
			"message": "Redirect Data",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
