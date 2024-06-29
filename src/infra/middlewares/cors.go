package middlewares

import (
	"net/http"
)

func EnableCORS(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	
	// Allow requests from any origin
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin,access-control-allow-headers,access-control-allow-methods,access-control-allow-credentials")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, HEAD ,GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	
	w.Write([]byte("Hello, World!"))
	next.ServeHTTP(w, r)
	
	})
	
	}
