package models

import (
	"time"
)

type User struct {
	UserID       uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	GoogleID     string    `gorm:"default:null" json:"google_id"`
}