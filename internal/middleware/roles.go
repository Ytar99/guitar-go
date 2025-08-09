package middleware

import (
	"net/http"

	"guitar-go/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

func (m *Middleware) RequireRole(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(ClaimsKey).(jwt.MapClaims)
			if !ok {
				utils.JSONError(w, http.StatusForbidden, "Invalid token claims")
				return
			}

			userRole, ok := claims["role"].(string)
			if !ok || userRole != requiredRole {
				utils.JSONError(w, http.StatusForbidden, "Insufficient permissions")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
