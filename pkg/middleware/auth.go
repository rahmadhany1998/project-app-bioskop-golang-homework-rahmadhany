package middleware

import (
	"context"
	"net/http"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/repository"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/utils"
	"strings"
)

type contextKey string

const (
	ContextUserID   contextKey = "user_id"
	ContextUsername contextKey = "username"
)

func AuthMiddlewareWithRepo(repo repository.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" || !strings.HasPrefix(token, "Bearer ") {
				utils.WriteError(w, "Missing or invalid token", http.StatusUnauthorized)
				return
			}
			token = strings.TrimPrefix(token, "Bearer ")
			valid, err := repo.TokenRepo.IsTokenValid(r.Context(), token)
			if err != nil || !valid {
				utils.WriteError(w, "Unauthorized or expired token", http.StatusUnauthorized)
				return
			}
			claims, err := utils.ValidateJWT(token)
			if err != nil {
				utils.WriteError(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), ContextUserID, claims.UserID)
			ctx = context.WithValue(ctx, ContextUsername, claims.Username)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
