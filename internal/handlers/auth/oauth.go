package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func getConfig() *oauth2.Config {
	clientID := viper.GetString("GOOGLE_CLIENT_ID")
	if clientID == "" {
		log.Fatalln("GOOGLE_CLIENT_ID is not specified.")
	}
	clientSecret := viper.GetString("GOOGLE_CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatalln("GOOGLE_CLIENT_SECRET is not specified.")
	}
	redirectURL := viper.GetString("GOOGLE_REDIRECT_URL")
	if redirectURL == "" {
		log.Fatalln("GOOGLE_REDIRECT_URL is not specified.")
	}
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

type Claims struct {
	UserID uuid.UUID
	jwt.RegisteredClaims
}


func (h *AuthHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := getConfig().AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

func (h *AuthHandler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Error: code is missing.", http.StatusBadRequest)
		return
	}

	token, err := getConfig().Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Error: Failed to exchange token", http.StatusBadRequest)
		return
	}

	client := getConfig().Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		http.Error(w, "Error: Failed to get user information", http.StatusBadRequest)
		return
	}

	defer resp.Body.Close()

	// TODO change error to proper json
	var userInfo struct {
		Name string `json:"name"`
		Sub  string `json:"sub"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Error: Failed to parse user information", http.StatusInternalServerError)
		return
	}

	var user models.User
	if err := h.db.Where("google_id = ?", userInfo.Sub).First(&user).Error; err != nil {
		if err := h.db.Create(&models.User{
			GoogleID:  userInfo.Sub,
			Username:  userInfo.Name,
			CreatedAt: time.Now(),
		}).Error; err != nil {
			http.Error(w, "Error: Failed to create user", http.StatusInternalServerError)
			return
		}
		if err := h.db.Where("google_id = ?", userInfo.Sub).First(&user).Error; err != nil {
			http.Error(w, "Error: Failed to fetch created user", http.StatusInternalServerError)
			return
		}
	}

	jwtSecret := viper.GetString("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatalln("JWT_SECRET is not specified.")
		http.Error(w, "Error: JWT_SECRET is not specified.", http.StatusInternalServerError)
		return
	}

	accessTokenClaims := Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "Error: Failed to generate access token.", http.StatusInternalServerError)
		return
	}

	refreshTokenClaims := Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "Error: Failed to generate refresh token.", http.StatusInternalServerError)
		return
	}

	// Instead of set the cookie for client, we let the client decide on how to store these tokens.
	type ReturnToken struct {
		RefreshToken string `json:"refresh_token"`
		AccessToken string `json:"access_token"`
		RefreshTokenMaxAge int `json:"refresh_token_max_age"`
		AccessTokenMaxAge int `json:"access_token_max_age"`
	}

	returnToken := ReturnToken {
		RefreshToken: refreshToken,
		AccessToken: accessToken,
		RefreshTokenMaxAge: 7 * 24 * 60 * 60, // 7 days
		AccessTokenMaxAge: 15 * 60, // 15 minutes
	}

	redirectURL := viper.GetString("FRONTEND_REDIRECT_URL")
	json.NewEncoder(w).Encode(returnToken)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}