package auth

import (
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Extract token from request header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" || !strings.Contains(tokenString, " ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString = strings.Split(tokenString, " ")[1]

		// Parse and validate the token
		claims, err := DecodeJwtToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "", "")

		// Add user claims to the request context
		ctx = context.WithValue(ctx, "user", claims)

		// Proceed to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
