package middleware

import (
	"context"
	"net/http"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
)

// random thing idk 
func UserCtx(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    user := models.User {
		Username: "1234",
	}
    ctx := context.WithValue(r.Context(), "user", user)
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}