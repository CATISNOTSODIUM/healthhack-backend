package models
import (
	"time"
	"gorm.io/gorm"
)
type VoiceActivityAnalysis struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Duration float32 `gorm:"not null" json:"duration"`
	SpeechDuration float32 `gorm:"not null" json:"speech_duration"`
	PausesDuration float32 `gorm:"not null" json:"pauses_duration"`
	NumberOfPauses uint `gorm:"not null" json:"number_of_pauses"`
	AnswerDelay float32 `gorm:"not null" json:"answer_delay"`
	// Detailed number of pauses
	Pauses []Pauses `json:"pauses"`
    DeletedAt time.Time
	IsDeleted gorm.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt" json:"deleted_at,omitempty"`
}
