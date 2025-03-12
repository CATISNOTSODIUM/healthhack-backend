package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/databases"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func main() {
	// Read environment variables
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err);
	}
	portNumber, ok := viper.Get("PORT").(string)
	if (!ok) {
		log.Println("PORT number not specified. Set to default value 8080.")
		portNumber = "8080"
	}
	
	// OpenAIToken is not necessary at this stage
	openAIToken, ok := viper.Get("OPENAI_TOKEN").(string)
	var openAIClient * openai.Client
	if (!ok) {
		log.Println("OPENAI_TOKEN not specified")
		openAIClient = nil
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