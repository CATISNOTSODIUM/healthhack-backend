package main

import (
	"log"
	"net/http"
	"os"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/databases"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		log.Println("PORT number not specified. Set to default value 8080.")
		portNumber = "8080"
	}
	db := databases.InitDB()
	r := chi.NewRouter()
	r.Group(routes.GetRoutes(db))

	log.Printf("Listening on port %s at http://localhost:${%s}!", portNumber, portNumber)
	log.Fatalln(http.ListenAndServe(":" + portNumber, r))
}