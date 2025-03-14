package textAnalysis

import (
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

type TextAnalysisHandler struct {
	db *gorm.DB
	openai *openai.Client
}

func NewTextAnalysisHandler(db *gorm.DB, openai *openai.Client) *TextAnalysisHandler {
	return &TextAnalysisHandler{db: db, openai: openai}
}