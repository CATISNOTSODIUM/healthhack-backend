package models

import "github.com/google/uuid"

type TextAnalysis struct {
	ID   uuid.UUID   	`gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	HistoryID uuid.UUID `gorm:"not null" json:"history_id"`
	TextSummary string `gorm:"not null" json:"text_summary"`
}