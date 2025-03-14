package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/middleware"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(middleware.ContextKey("user_id")).(uuid.UUID)
	if (!ok) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid context. userID should be type uuid."))
		return 
	}
	
	user := models.User{}
	result := h.db.Clauses(clause.Returning{}).First(&user, "id = ?", userID)
	if result.Error != nil {
		if (errors.Is(result.Error, gorm.ErrRecordNotFound)) {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("User %s not founded", userID)))
        	return 
		}
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
        return 
	}
	returnResponse := struct {
		Username string
		Age uint
		MedicalRecord string
	} {
		Username: user.Username,
		Age: user.Age,
		MedicalRecord: user.MedicalRecord,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnResponse)
}