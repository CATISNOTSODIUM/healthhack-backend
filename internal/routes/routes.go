package routes

import (
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/handlers/history"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/handlers/user"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/handlers/voiceAnalysis"
	"github.com/go-chi/chi/v5"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

func GetRoutes(db *gorm.DB, openAIClient *openai.Client) func(r chi.Router) {
	voiceAnalysisHandler := voiceAnalysis.NewVoiceAnalysisHandler(db)
	userHandler := user.NewUserHandler(db)
	historyHandler := history.NewHistoryHandler(db)
	return func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to the server"))
		})
		r.Route("/api", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/get", userHandler.GetUser)
				r.Post("/create", userHandler.CreateUser)
				r.Put("/update", userHandler.UpdateUser)
				r.Put("/delete", userHandler.DeleteUser)
			})
			
			r.Route("/history", func(r chi.Router) {
				r.Post("/create", historyHandler.CreateHistory)
				r.Get("/get", historyHandler.GetUserHistories)
			})
		})
		r.Route("/internal", func(r chi.Router) {
			r.Put("/voice-analysis/create", voiceAnalysisHandler.CreateRecordFromHistoryID)
		})
		
	}
}
