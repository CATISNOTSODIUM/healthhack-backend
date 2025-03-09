package history

import (
	"fmt"
	"net/http"
	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
)


func (h *HistoryHandler) GetUserHistory(w http.ResponseWriter, r *http.Request) {
	// TODO
	ctx := r.Context()
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	var err = h.db.First(&user, "username = ?", user.Username).Error
	if err != nil {
		w.Write([]byte(fmt.Sprintf("User %s not found", user.Username)))
		return
	}

	w.Write([]byte(fmt.Sprintf("User:%s", user.Username)))
}
