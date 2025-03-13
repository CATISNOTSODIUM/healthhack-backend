package textAnalysis

import (
	"encoding/json"
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/api"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)


func (h *TextAnalysisHandler) CreateTextRecord(w http.ResponseWriter, r *http.Request) {
	type Request struct{
		HistoryID uuid.UUID `json:"history_id"`
		TranscribedText string `json:"transcribed_text"`
	} 

	request := &Request{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return 
	}

	textSummary, err := api.ExtractTextFeature(request.TranscribedText, h.openai)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write response back to database
	newRecord := models.TextAnalysis {
		HistoryID: request.HistoryID,
		TextSummary: textSummary,
	}
	
	result := h.db.Clauses(clause.Returning{}).Create(&newRecord)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return 
	}

	json.NewEncoder(w).Encode(map[string]string{
		"response": textSummary,
		"error": result.Error.Error(),
	})

	w.WriteHeader(http.StatusInternalServerError)
}