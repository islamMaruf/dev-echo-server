package middleware

import "net/http"

// Security returns a middleware that adds security headers to HTTP responses.
// Provides protection similar to Helmet.js for Node.js applications.
//
// Headers added:
//   - X-Content-Type-Options: nosniff - Prevents MIME type sniffing
//   - X-Frame-Options: DENY - Prevents clickjacking attacks
//   - X-XSS-Protection: 1; mode=block - Enables XSS filtering
//   - Strict-Transport-Security: Forces HTTPS connections (31536000 = 1 year)
//   - Content-Security-Policy: default-src 'self' - Restricts resource loading
//
// Parameters:
//   - next: The next handler in the middleware chain
//
// Returns:
//   - http.Handler: Handler with security headers applied
func Security(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set security headers
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")

		next.ServeHTTP(w, r)
	})
}
