package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/middleware"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

func (h *AuthHandler) HandleRefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization") // refresh token
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Authorization header is missing",
		})
		return
	}
	tokenString = tokenString[len("Bearer "):]
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatalln("JWT_SECRET is not specified.")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "JWT_SECRET is not specified.",
		})
		return
	}
	refreshClaims := &middleware.Claims{}
	refreshTokenParsed, err := jwt.ParseWithClaims(tokenString, refreshClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !refreshTokenParsed.Valid || refreshClaims.ExpiresAt.Before(time.Now()) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User not found",
		})
		return
	}

	var user models.User
	if err := h.db.First(&user, refreshClaims.UserID).Error; err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User not found",
		})
		return
	}

	accessTokenClaims := middleware.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}
	newAccessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte(jwtSecret))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to generate new access token",
		})
		return
	}

	type ReturnToken struct {
		AccessToken string `json:"access_token"`
	}

	returnToken := ReturnToken {
		AccessToken: newAccessToken,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnToken)
}
