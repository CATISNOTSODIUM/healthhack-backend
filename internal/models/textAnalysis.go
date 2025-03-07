package models

type TextAnalysis struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Summary string `gorm:"not null" json:"summary"`
}