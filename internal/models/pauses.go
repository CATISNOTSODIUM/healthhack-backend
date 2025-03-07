package models

import "time"

type Pauses struct {
    ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    VoiceActivityAnalysisID uint `gorm:"not null" json:"voice_activity_analysis_id"`
    StartTime time.Time `gorm:"not null" json:"start_time"`
    EndTime   time.Time `gorm:"not null" json:"end_time"`
    Duration  float32    `gorm:"not null" json:"duration"` 
}