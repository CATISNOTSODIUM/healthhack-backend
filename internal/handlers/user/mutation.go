package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Username string `json:"username" validate:"required"`
		GoogleID string `json:"google_id" validate:"required"`
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
	
	newUser := models.User {
		Username: request.Username,
		GoogleID: request.GoogleID,
	}
	
	result := h.db.Clauses(clause.Returning{}).Create(&newUser)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
        return 
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(newUser)
}


func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		ID uuid.UUID `json:"id" validate:"required"`
		Username string `json:"username"`
		Age uint `json:"age"`
		MedicalRecord string `json:"medical_record"`
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

	newUser := models.User {
		Username: request.Username,
		Age: request.Age,
		MedicalRecord: request.MedicalRecord,
	}
	
	result := h.db.Clauses(clause.Returning{}).Save(&newUser)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
        return 
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(newUser)
}


func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
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
	
	result := h.db.Delete(&models.User{}, request.ID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
        return 
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(fmt.Sprintf("Successfully delete user id:%d", request.ID)))
}
