package history

import (
	"encoding/json"
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/middleware"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func (h *HistoryHandler) CreateHistory(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	userID, ok := ctx.Value(middleware.ContextKey("user_id")).(uuid.UUID)
	if (!ok) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid context. userID should be type uuid."))
		return 
	}

	newHistory := models.History {
		UserID: userID,
	}
	
	result := h.db.Clauses(clause.Returning{}).Create(&newHistory)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
        return 
	}

	response := models.History {
		ID: newHistory.ID,
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}


func (h *HistoryHandler) DeleteHistoryByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(middleware.ContextKey("user_id")).(uuid.UUID)
	if (!ok) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid context. userID should be type uuid."))
		return 
	}

	type Request struct {
		HistoryID uuid.UUID `json:"id" validate:"required"`
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
	
	newHistory := models.History {
		UserID: userID,
		ID: request.HistoryID,
	}
	
	result := h.db.Clauses(clause.Returning{}).Delete(&newHistory)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
        return 
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Successfully delete history."))
}

