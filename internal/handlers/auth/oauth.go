package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/middleware"
	"github.com/golang-jwt/jwt/v5"
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



func (h *AuthHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := getConfig().AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

func (h *AuthHandler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "code is missing",
		})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := getConfig().Exchange(context.Background(), code)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed exchange token",
		})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := getConfig().Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to get user information",
		})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var userInfo struct {
		Name string `json:"name"`
		Sub  string `json:"sub"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to parse user information",
		})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user models.User
	if err := h.db.Where("google_id = ?", userInfo.Sub).First(&user).Error; err != nil {
		if err := h.db.Create(&models.User{
			GoogleID:  userInfo.Sub,
			Username:  userInfo.Name,
			CreatedAt: time.Now(),
		}).Error; err != nil {
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Failed to create user",
			})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := h.db.Where("google_id = ?", userInfo.Sub).First(&user).Error; err != nil {
			json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to fetch created user",
			})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	jwtSecret := viper.GetString("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatalln("JWT_SECRET is not specified.")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "JWT_SECRET is not specified.",
		})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accessTokenClaims := middleware.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte(jwtSecret))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to generate access token.",
		})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	refreshTokenClaims := middleware.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString([]byte(jwtSecret))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to generate refresh token.",
		})
		w.WriteHeader(http.StatusInternalServerError)
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
		AccessTokenMaxAge: 150 * 60, // 150 minutes
	}

	redirectURL := viper.GetString("FRONTEND_REDIRECT_URL")
	json.NewEncoder(w).Encode(returnToken)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}