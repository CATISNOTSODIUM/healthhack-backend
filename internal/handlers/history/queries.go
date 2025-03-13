package history

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/middleware"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (h *HistoryHandler) GetUserHistories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(middleware.ContextKey("user_id")).(uuid.UUID)
	if (!ok) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid context. userID should be type uuid."))
		return 
	}

	// get latest histories
	var retrievedHistories []models.History
	result := h.db.Where(&models.History{UserID: userID}).Find(&retrievedHistories)
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

func (h *HistoryHandler) GetLatestHistory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(middleware.ContextKey("user_id")).(uuid.UUID)
	if (!ok) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid context. userID should be type uuid."))
		return 
	}
	var retrievedHistory models.History
	result := h.db.Order("created_at DESC").Where(&models.History{UserID: userID}).First(&retrievedHistory)
	if result.Error != nil {
		w.Write([]byte("cannot retrieve latest history"))
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(retrievedHistory)
}

func (h *HistoryHandler) GetUserHistoryByID(w http.ResponseWriter, r *http.Request) {
	historyID := r.URL.Query().Get("id") 
	log.Println("history id", historyID)
	// get latest histories
	var retrievedHistory models.History
	result := h.db.Clauses(clause.Returning{}).First(&retrievedHistory, "id = ?", historyID)
	if result.Error != nil {
		if (errors.Is(result.Error, gorm.ErrRecordNotFound)) {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("History %s not founded", historyID)))
        	return 
		}
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(retrievedHistory)
}
