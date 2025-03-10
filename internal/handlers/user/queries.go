package user

import (
	"encoding/json"
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"gorm.io/gorm/clause"
)


func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id") 
	var user models.User
	result := h.db.Clauses(clause.Returning{}).Find(&user, userID)
	if result.Error != nil {
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