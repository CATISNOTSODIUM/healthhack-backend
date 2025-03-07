package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/databases"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/routes"
	"github.com/go-chi/chi/v5"
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

	db := databases.InitDB()
	r := chi.NewRouter()
	r.Group(routes.GetRoutes(db))
	
	log.Println("Listening on port " + portNumber)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", portNumber), r))
}