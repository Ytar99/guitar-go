package middleware

import (
	"context"
	"net/http"
	"strings"

	"guitar-go/internal/services"
	"guitar-go/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

type Middleware struct {
	authService services.AuthService
}

const ClaimsKey contextKey = "claims"

func NewMiddleware(authSvc services.AuthService) *Middleware {
	return &Middleware{authService: authSvc}
}

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.JSONError(w, http.StatusUnauthorized, "Authorization header required")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := m.authService.ParseToken(tokenString)
		if err != nil || !token.Valid {
			utils.JSONError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.JSONError(w, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		ctx := context.WithValue(r.Context(), ClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
