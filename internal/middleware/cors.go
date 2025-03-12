package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/cors"
	"github.com/spf13/viper"
)

func CorsMiddleware() func(next http.Handler) http.Handler {
	allowedOriginsEnv := viper.GetString("ALLOWED_ORIGINS")
	if allowedOriginsEnv == "" {
		log.Println("Warning: ALLOWED_ORIGINS not set, using default '*'")
		allowedOriginsEnv = "*"
	}

	allowedOrigins := strings.Split(allowedOriginsEnv, ",")
	return cors.Handler(cors.Options{
		AllowedOrigins:     allowedOrigins,
		AllowedMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposedHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}

func CorsMiddlewareInternal() func(next http.Handler) http.Handler {
	allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS_INTERNAL")
	if allowedOriginsEnv == "" {
		log.Println("Warning: ALLOWED_ORIGINS not set, using default '*'")
		allowedOriginsEnv = "*"
	}

	allowedOrigins := strings.Split(allowedOriginsEnv, ",")
	return cors.Handler(cors.Options{
		AllowedOrigins:     allowedOrigins,
		AllowedMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposedHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}