package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/juaniferro/fake-twitter/internal/models"
	"github.com/juaniferro/fake-twitter/internal/usecases"
)

type GetTimelineHandlerInterface interface{
	TimelineGetter(user int) (tweets []models.TimelineTweet, err error)
}

type GetTimelineHandler struct {
	getTimelineUsecase usecases.GetTimelineUsecase
}

func NewGetTimelineHandler(getTimelineUsecase usecases.GetTimelineUsecase) *GetTimelineHandler{
	return &GetTimelineHandler{getTimelineUsecase: getTimelineUsecase}
}

func (gth GetTimelineHandler)HandleGetTimeline(w http.ResponseWriter, r *http.Request) {
	user_id_str := r.Header.Get("user_id")
	user_id, err := strconv.Atoi(user_id_str) 
	if err != nil {
        http.Error(w, "Invalid user_id", http.StatusBadRequest)
        return
    }

	timeline, err := gth.getTimelineUsecase.TimelineGetter(user_id)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(timeline)
}