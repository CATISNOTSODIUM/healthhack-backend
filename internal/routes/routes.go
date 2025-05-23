package routes

import (
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/handlers/auth"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/handlers/history"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/handlers/textAnalysis"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/handlers/user"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/handlers/voiceAnalysis"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
	OpenAIClient *openai.Client
}

func GetRoutes(config Config) func(r chi.Router) {
	voiceAnalysisHandler := voiceAnalysis.NewVoiceAnalysisHandler(config.DB)
	userHandler := user.NewUserHandler(config.DB)
	historyHandler := history.NewHistoryHandler(config.DB)
	authHandler := auth.NewAuthHandler(config.DB)
	textAnalysisHandler := textAnalysis.NewTextAnalysisHandler(config.DB, config.OpenAIClient)
	
	return func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to the server (V2)"))
		})
		r.Route("/api", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(config.DB))
			r.Route("/users", func(r chi.Router) {
				r.Get("/get", userHandler.GetUser)
				r.Put("/update", userHandler.UpdateUser)
				r.Put("/delete", userHandler.DeleteUser)
			})
			r.Route("/history", func(r chi.Router) {
				r.Post("/create", historyHandler.CreateHistory)
				r.Get("/get", historyHandler.GetUserHistories)
				r.Get("/get-latest", historyHandler.GetLatestHistory)
				r.Get("/get-by-id", historyHandler.GetUserHistoryByID)
				r.Post("/create-text-analysis", textAnalysisHandler.CreateTextRecord)
			})
		})
		r.Route("/api/auth/google", func(r chi.Router) {
			r.Get("/", authHandler.GoogleLogin)
			r.Get("/callback", authHandler.GoogleCallback)
			r.Get("/refresh", authHandler.HandleRefreshToken) 
		})
		r.Route("/internal", func(r chi.Router) {
			r.Post("/voice-analysis/create", voiceAnalysisHandler.CreateRecordFromHistoryID)
		})
	}
}
