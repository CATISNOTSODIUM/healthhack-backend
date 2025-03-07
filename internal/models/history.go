package models

import "time"

type History struct {
	ID                      uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID                  uint      `gorm:"not null" json:"user_id"`
	VoiceActivityAnalysisID uint      `gorm:"not null" json:"voice_activity_analysis_id"`
	TextAnalysisID          uint      `gorm:"not null" json:"text_analysis_id"`
	CreatedAt               time.Time `gorm:"autoCreateTime" json:"created_at"`
}
