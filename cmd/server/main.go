package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/databases"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/sashabaranov/go-openai"
	"github.com/joho/godotenv"
)

func main() {
	// Read environment variables
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("err loading: %v", err)
    }

	portNumber := os.Getenv("PORT")
	if (portNumber == "") {
		log.Println("PORT number not specified. Set to default value 8080.")
		portNumber = "8080"
	}
	
	// OpenAIToken is not necessary at this stage
	openAIToken := os.Getenv("OPENAI_TOKEN")
	var openAIClient * openai.Client
	if (openAIToken == "") {
		log.Println("OPENAI_TOKEN not specified")
	} else {
		openAIClient = openai.NewClient(openAIToken)
	}

	// Init instances
	db := databases.InitDB()
	// Init router
	r := chi.NewRouter()
	r.Group(routes.GetRoutes(routes.Config{DB: db, OpenAIClient: openAIClient}))
	
	log.Println("Listening on port " + portNumber)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", portNumber), r))
}