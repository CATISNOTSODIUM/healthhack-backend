package models

import "github.com/google/uuid"

type VoiceActivityAnalysis struct {
	ID   uuid.UUID   	`gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	HistoryID uuid.UUID `gorm:"not null" json:"history_id"`
	Duration float32 `gorm:"not null" json:"duration"`
	TotalSpeechDuration float32 `gorm:"not null" json:"total_speech_duration"`
	TotalPausesDuration float32 `gorm:"not null" json:"total_pauses_duration"`
	NumSpeechSegments uint `gorm:"not null" json:"num_speech_segments"`
	NumPauses uint `gorm:"not null" json:"num_pauses"`
	AnswerDelayDuration float32 `gorm:"not null" json:"answer_delay_duration"`
	
	Pauses []Pause `gorm:"ForeignKey:VoiceActivityAnalysisID" json:"pauses"`
	SpeechSegments []SpeechSegment `gorm:"ForeignKey:VoiceActivityAnalysisID" json:"speech_segments"`
	// Voice file is deleted immediately after the analysis is completed.
}
