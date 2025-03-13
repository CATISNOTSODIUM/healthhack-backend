package voiceAnalysis

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/healthhack-backend/internal/models"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (h *VoiceAnalysisHandler) CreateRecordFromHistoryID(w http.ResponseWriter, r *http.Request) {
	request := &VoiceActivityAnalysisRequest{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return 
	}

	if err := validator.New().Struct(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return 
	}

	// check if historyID exists
	var exists bool
	err := h.db.Model(models.History{}).Select("count(*) > 0").Where("id = ?", request.HistoryID).Find(&exists).Error
	if err != nil {
		http.Error(w, "historyID does not exist.", http.StatusBadRequest)
        return 
	}

	// create voice analysis request
	analysis := models.VoiceActivityAnalysis {
		HistoryID:  request.HistoryID,
		Duration: request.Duration,
		TotalSpeechDuration: request.TotalSpeechDuration,
		TotalPausesDuration: request.TotalSpeechDuration,
		NumSpeechSegments: request.NumSpeechSegments,
		NumPauses: request.NumPauses,
		AnswerDelayDuration: request.AnswerDelayDuration,
	}

	tmp := models.VoiceActivityAnalysis {}
	if err := h.db.Clauses(clause.Returning{}).Where("history_id = ?", request.HistoryID).First(&tmp).Error; err != nil {
		if (errors.Is(err, gorm.ErrRecordNotFound)) {
			h.db.Clauses(clause.Returning{}).Create(&analysis)
		} 
	} else {
		result := h.db.Clauses(clause.Returning{}).Where("history_id = ?", request.HistoryID).Updates(&analysis)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return 
		}
	}
	
	voiceActivityAnalysisID := analysis.ID

	for _, pause := range request.Pauses {
		newPause := models.Pause{
			StartTime: pause.StartTime,
			EndTime: pause.EndTime,
			Duration: pause.Duration,
			VoiceActivityAnalysisID: voiceActivityAnalysisID,
		}
		result := h.db.Create(&newPause)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			return 
		}
	}

	for _, segments := range request.SpeechSegments {
		newSegment := models.Pause{
			StartTime: segments.StartTime,
			EndTime: segments.EndTime,
			Duration: segments.Duration,
			VoiceActivityAnalysisID: voiceActivityAnalysisID,
		}
		result := h.db.Create(&newSegment)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			return 
		}
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(fmt.Sprintf("Successfully update history %s with voice activity analysis %s",  request.HistoryID, voiceActivityAnalysisID)))
}
