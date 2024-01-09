package server

import (
	"net/http"
	"strings"
	"vg-sage/internal/utils"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		validToken := utils.ValidateToken(tokenString)
		if !validToken {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
