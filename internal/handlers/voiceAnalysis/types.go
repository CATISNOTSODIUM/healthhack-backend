package voiceAnalysis

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoiceAnalysisHandler struct {
	db *gorm.DB
}

func NewVoiceAnalysisHandler(db *gorm.DB) *VoiceAnalysisHandler {
	return &VoiceAnalysisHandler{db: db}
}

type VoiceActivityAnalysisRequest struct {
	HistoryID uuid.UUID `json:"history_id" validate:"required"`
	Duration float32 `json:"total_duration"`
	TotalSpeechDuration float32 `json:"total_speech_duration"`
	TotalPausesDuration float32 `json:"total_pause_duration"`
	NumSpeechSegments uint `json:"num_speech_segments"`
	NumPauses uint `json:"num_pauses"`
	AnswerDelayDuration float32 `json:"answer_delay_duration"`
	Pauses []PausesRequest `json:"pause_segments"`
	SpeechSegments []SpeechSegmentsRequest `json:"speech_segments"`
}

type PausesRequest struct {
    StartTime float32 `gorm:"not null" json:"start_time"`
    EndTime   float32 `gorm:"not null" json:"end_time"`
    Duration  float32    `gorm:"not null" json:"duration"` 
}

type SpeechSegmentsRequest struct {
    StartTime float32 `gorm:"not null" json:"start_time"`
    EndTime   float32 `gorm:"not null" json:"end_time"`
    Duration  float32    `gorm:"not null" json:"duration"` 
}