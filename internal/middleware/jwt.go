package middleware

import (
	"auth/internal/utils"
	"net/http"
	"strings"
)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized, please login first", http.StatusUnauthorized) // return warning that user need to login
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		r.Header.Set("X-User-ID", string(rune(claims.UserID)))
		r.Header.Set("X-Username", claims.Username)

		next(w, r)
	}
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-Username")

	w.Write([]byte("Hello, " + username)) // return greeting to user
}
