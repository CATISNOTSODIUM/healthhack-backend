package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/middleware"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

/*
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Username string `json:"username" validate:"required"`
		GoogleID string `json:"google_id" validate:"required"`
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
		GoogleID: request.GoogleID,

	}

	result := h.db.Clauses(clause.Returning{}).Create(&newUser)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUser)
}
*/

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(middleware.ContextKey("user_id")).(uuid.UUID)
	if (!ok) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid context. userID should be type uuid."))
		return 
	}

	type Request struct {
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
		ID: userID,
		Username: request.Username,
		Age: request.Age,
		MedicalRecord: request.MedicalRecord,
	}
	
	result := h.db.Clauses(clause.Returning{}).Save(&newUser)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return 
	}

	returnResponse := struct {
		Username string
		Age uint
		MedicalRecord string
	} {
		Username: newUser.Username,
		Age: newUser.Age,
		MedicalRecord: newUser.MedicalRecord,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnResponse)
}


func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(middleware.ContextKey("user_id")).(uuid.UUID)
	if (!ok) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid context. userID should be type uuid."))
		return 
	}

	result := h.db.Delete(&models.User{}, userID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return 
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Successfully delete user id:%s", userID)))
}
