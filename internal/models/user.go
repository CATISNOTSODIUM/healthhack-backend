package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID   uuid.UUID   	`gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Username string `gorm:"not null" json:"username"`
	GoogleID string `gorm:"default:null" json:"google_id"`

	// Personal information
	Age uint `gorm:"default:0" json:"age"`
	MedicalRecord string `gorm:"default:null" json:"medical_record"`

	// Analysis histories
	Histories []History `json:"histories"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt time.Time
	IsDeleted gorm.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt" json:"deleted_at,omitempty"`
}
