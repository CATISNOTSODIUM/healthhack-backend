package history

import (
	"encoding/json"
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func (h *HistoryHandler) CreateHistory(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		UserID uuid.UUID `json:"user_id" validate:"required"`
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
		UserID: request.UserID,
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

