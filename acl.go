package acl

import "net/http"

// Middleware function to restrict access to certain endpoints
func restrictEndpoints(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is to a restricted endpoint
		if isRestrictedEndpoint(r.URL.Path) {
			http.Error(w, "Access Forbidden", http.StatusForbidden)
			return
		}

		// If not restricted, call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// Function to check if an endpoint is restricted
func isRestrictedEndpoint(endpoint string) bool {
	restrictedEndpoints := []string{"/restricted", "/admin"}

	for _, e := range restrictedEndpoints {
		if e == endpoint {
			return true
		}
	}

	return false
}
