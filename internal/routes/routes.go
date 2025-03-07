package routes

import (
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/handlers/users"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/middleware"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func GetRoutes(db *gorm.DB) func(r chi.Router) {
    userHandler := users.NewUserHandler(db)
   
    return func(r chi.Router) {
        r.Get("/", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("welcome to the server"))
        })
        r.Route("/users", func(r chi.Router) {
            r.Use(middleware.UserCtx)
            r.Get("/", userHandler.GetUser)
        })
    }
}



