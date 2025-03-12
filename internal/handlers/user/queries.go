package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		ID uuid.UUID `json:"id" validate:"required"`
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
	
	user := models.User{}
	result := h.db.Clauses(clause.Returning{}).First(&user, "id = ?", request.ID)
	if result.Error != nil {
		if (errors.Is(result.Error, gorm.ErrRecordNotFound)) {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("User %s not founded", request.ID)))
        	return 
		}
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
        return 
	}
	returnResponse := models.User{
		ID: user.ID,
		Username: user.Username,
		Age: user.Age,
		MedicalRecord: user.MedicalRecord,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnResponse)
}