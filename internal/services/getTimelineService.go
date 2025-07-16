package services

import (
	"github.com/juaniferro/fake-twitter/internal/models"
	"github.com/juaniferro/fake-twitter/internal/repositories"
)

type GetTimelineService struct {
	fakeTwitterRepo repositories.FakeTwitterRepo
}

type GetTimelineServiceServiceInterface interface{
	GetTimeline(user int) (tweets []models.TimelineTweet, err error)
}

func NewGetTimelineService(fakeTwitterRepo repositories.FakeTwitterRepo) *GetTimelineService{
	return &GetTimelineService{fakeTwitterRepo: fakeTwitterRepo}
}

func (gts GetTimelineService) GetTimelineCaller(user int) (tweets []models.TimelineTweet, err error) {
	return gts.fakeTwitterRepo.GetTimeline(user)
}
