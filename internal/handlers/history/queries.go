package history

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/middleware"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO: use context
func (h *HistoryHandler) GetUserHistories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(middleware.ContextKey("user_id")).(uuid.UUID)
	if (!ok) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid context. userID should be type uuid."))
		return 
	}

	NumberOfHistories, err := strconv.Atoi(r.URL.Query().Get("number_of_histories"))
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: number_of_histories should be unsigned integer"))
		return 
    }

	// get latest histories
	var retrievedHistories []models.History
	result := h.db.Where(&models.History{UserID: userID}).Limit(NumberOfHistories).Find(&retrievedHistories)
	if result.Error != nil {
		if (errors.Is(result.Error, gorm.ErrRecordNotFound)) {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("User %s not founded", userID)))
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
