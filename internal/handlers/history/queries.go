package history

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *HistoryHandler) GetUserHistories(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		UserID            uuid.UUID `json:"user_id" validate:"required"`
		NumberOfHistories int       `json:"number_of_histories"`
	}
	request := &Request{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.New().Struct(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get latest histories
	var retrievedHistories []models.History
	result := h.db.Where(&models.History{UserID: request.UserID}).Limit(request.NumberOfHistories).Find(&retrievedHistories)
	if result.Error != nil {
		if (errors.Is(result.Error, gorm.ErrRecordNotFound)) {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("User %s not founded", request.UserID)))
        	return 
		}
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	for i, history := range retrievedHistories {
		var retrievedVoiceAnalysis models.VoiceActivityAnalysis
		var retrievedTextAnalysis models.TextAnalysis
		result := h.db.Where(&models.VoiceActivityAnalysis{HistoryID: history.ID}).Find(&retrievedVoiceAnalysis)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		result = h.db.Where(&models.TextAnalysis{HistoryID: history.ID}).Find(&retrievedTextAnalysis)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		retrievedHistories[i].VoiceActivityAnalysis = retrievedVoiceAnalysis
		retrievedHistories[i].TextAnalysis = retrievedTextAnalysis
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(retrievedHistories)
}
