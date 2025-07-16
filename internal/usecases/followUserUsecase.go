package usecases

import (

	"github.com/juaniferro/fake-twitter/internal/services"
)

type FollowUserUsecase struct {
	followUserService services.FollowUserService
}

type FollowUserUsecaseInterface interface{
	FollowUserCaller(user, followedUser int) error 
}

func NewFollowUsertUsecase(followUserService services.FollowUserService) *FollowUserUsecase {
	return &FollowUserUsecase{followUserService: followUserService}
}

func (fuu FollowUserUsecase) UserFollower(user, followedUser int) error {
	return fuu.followUserService.FollowUserCaller(user, followedUser)
}