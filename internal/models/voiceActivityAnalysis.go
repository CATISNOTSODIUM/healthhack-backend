package models


type VoiceActivityAnalysis struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	DownloadURL string `gorm:"not null" json:"download_url"`
	Duration float32 `gorm:"not null" json:"duration"`
	SpeechDuration float32 `gorm:"not null" json:"speech_duration"`
	PausesDuration float32 `gorm:"not null" json:"pauses_duration"`
	NumberOfPauses uint `gorm:"not null" json:"number_of_pauses"`
	AnswerDelay float32 `gorm:"not null" json:"answer_delay"`
}