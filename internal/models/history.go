package models

import (
	"time"
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	ID                      uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID                  uint      `gorm:"not null" json:"user_id"`
	VoiceActivityAnalysisID uint      `gorm:"not null" json:"voice_activity_analysis_id"`
	TextAnalysisID          uint      `gorm:"not null" json:"text_analysis_id"`
	CreatedAt               time.Time `gorm:"autoCreateTime" json:"created_at"`
}
