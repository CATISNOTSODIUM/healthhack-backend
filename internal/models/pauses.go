package models

import "time"

type Pauses struct {
    ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    StartTime time.Time `gorm:"default:null" json:"start_time"`
    EndTime   time.Time `gorm:"default:null" json:"end_time"`
    Duration  float32    `gorm:"default:0" json:"duration"` 
}