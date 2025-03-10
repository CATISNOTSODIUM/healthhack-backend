package models

import "github.com/google/uuid"
type Pause struct {
    ID   uuid.UUID   	`gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
    VoiceActivityAnalysisID uuid.UUID `gorm:"not null" json:"voice_activity_analysis_id"`
    StartTime float32 `gorm:"not null" json:"start_time"`
    EndTime   float32 `gorm:"not null" json:"end_time"`
    Duration  float32    `gorm:"not null" json:"duration"` 
}