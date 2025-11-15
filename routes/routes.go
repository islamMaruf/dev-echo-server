package routes

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterRoutes registers all application routes
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
