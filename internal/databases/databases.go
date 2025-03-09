package databases

import (
	"log"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	dsn, ok := viper.Get("DSN").(string)

	if (!ok) {
		log.Fatalf("DSN is not set in the environment variables")
	}

	log.Println("Connect to database")
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.History{},
		&models.TextAnalysis{},
		&models.VoiceActivityAnalysis{},
		&models.Pause{},
		&models.SpeechSegment{},
		&models.TextAnalysis{},
	); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	log.Println("Database connected, migrated, and categories added successfully")

	return db
}
