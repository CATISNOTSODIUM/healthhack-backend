package models


type VoiceActivityAnalysis struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	DownloadURL string `gorm:"default:null" json:"download_url"`
	Duration float32 `gorm:"default:0" json:"duration"`
	SpeechDuration float32 `gorm:"default:0" json:"speech_duration"`
	PausesDuration float32 `gorm:"default:0" json:"pauses_duration"`
	NumberOfPauses uint `gorm:"default:0" json:"number_of_pauses"`
	AnswerDelay float32 `gorm:"default:0" json:"answer_delay"`
}