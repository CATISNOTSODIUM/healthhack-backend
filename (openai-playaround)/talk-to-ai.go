package main

import (
	"log"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/api"
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

	openAIToken, ok := viper.Get("OPENAI_TOKEN").(string)
	if (!ok) {
		log.Fatalf("OPENAI_TOKEN not specified")
	}

	OBAMA_TEXT := "I stand here today humbled by the task before us, grateful for the trust you have bestowed, mindful of the sacrifices borne by our ancestors. I thank President Bush for his service to our nation, as well as the generosity and cooperation he has shown throughout this transition. Forty-four Americans have now taken the presidential oath. The words have been spoken during rising tides of prosperity and the still waters of peace. Yet, every so often the oath is taken amidst gathering clouds and raging storms. At these moments, America has carried on not simply because of the skill or vision of those in high office, but because We the People have remained faithful to the ideals of our forbearers, and true to our founding documents."
	openAIClient := openai.NewClient(openAIToken)
	response, _ := api.ExtractTextFeature(OBAMA_TEXT, openAIClient)
	log.Println("test OPENAI " + response)
}