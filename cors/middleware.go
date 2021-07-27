package cors

import "net/http"

// Middleware :
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allowed-Origin", "*")
		w.Header().Add("Content-type", "application/json")
		w.Header().Add("Access-Control-Allowed-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept")

		next.ServeHTTP(w, r)
	})
}
