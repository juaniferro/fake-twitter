package usecases

import (
	"github.com/juaniferro/fake-twitter/internal/models"
	"github.com/juaniferro/fake-twitter/internal/services"
)

type GetTimelineUsecase struct {
	GetTimelineService GetTimelineServiceInterface
}

type GetTimelineServiceInterface interface{
	GetTimelineCaller(user int) (tweets []models.TimelineTweet, err error) 
}

func NewGetTimelineUsecase(GetTimelineService services.GetTimelineService) *GetTimelineUsecase {
	return &GetTimelineUsecase{GetTimelineService: GetTimelineService}
}

func (gtu GetTimelineUsecase) TimelineGetter(user int) (tweets []models.TimelineTweet, err error) {
	return gtu.GetTimelineService.GetTimelineCaller(user)
}