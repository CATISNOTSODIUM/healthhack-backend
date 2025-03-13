package models

import "github.com/google/uuid"

type TextAnalysis struct {
	ID   uuid.UUID   	`gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	HistoryID uuid.UUID `gorm:"not null" json:"history_id"`
	CoherenceScore uint `gorm:"not null" json:"coherence_score"`
	CoherenceDescription string `gorm:"not null" json:"coherence_description"`
	SentenceComplexityScore uint `gorm:"not null" json:"sentence_complexity_score"`
	SentenceComplexityDescription string `gorm:"not null" json:"sentence_complexity_description"`
}