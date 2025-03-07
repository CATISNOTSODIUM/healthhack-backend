package models

type TextAnalysis struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Summary string `gorm:"default:null" json:"summary"`
}