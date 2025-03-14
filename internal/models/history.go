package models

import (
	"time"

	"github.com/google/uuid"
)

type History struct {
	ID                    uuid.UUID             `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID                uuid.UUID             `gorm:"not null" json:"user_id"`
	VoiceActivityAnalysis VoiceActivityAnalysis `gorm:"default:null;foreignKey:HistoryID" json:"voice_activity_analysis"`
	TextAnalysis          TextAnalysis          `gorm:"default:null;foreignKey:HistoryID" json:"text_analysis"`
	CreatedAt             time.Time             `gorm:"autoCreateTime" json:"created_at"`
}
