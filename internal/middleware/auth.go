package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Claims struct {
	UserID uuid.UUID
	jwt.RegisteredClaims
}

// TODO: add /refresh to add refresh token
func AuthMiddleware(db *gorm.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Authorization header is missing",
				})
				return
			}
			tokenString = tokenString[len("Bearer "):]
			jwtSecret := viper.GetString("JWT_SECRET")
			if jwtSecret == "" {
				log.Fatalln("JWT_SECRET is not specified.")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "JWT_SECRET is not specified.",
				})
				return
			}

			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			})
			
			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
						"error": "Invalid access token. You can refresh your token from /refresh",
				})
				return
			}

			var user models.User
			if err := db.First(&user, claims.UserID).Error; err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "User not found",
				})
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
