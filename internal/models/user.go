package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	GoogleID string `gorm:"default:null" json:"google_id"`

	Age uint `gorm:"default:0" json:"age"`
	// To be fixed
	MedicalRecord string `gorm:"default:null" json:"medical_record"`

	// Analysis histories
	Histories []History `json:"histories"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt time.Time
	IsDeleted gorm.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt" json:"deleted_at,omitempty"`
}
