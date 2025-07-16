package services

import "github.com/juaniferro/fake-twitter/internal/repositories"

type FollowUserService struct {
	fakeTwitterRepo FollowUserRepositoryInterface
}

type FollowUserRepositoryInterface interface{
	FollowUser(user, followedUser int) error
}

func NewFollowUserService(fakeTwitterRepo repositories.FakeTwitterRepo) *FollowUserService{
	return &FollowUserService{fakeTwitterRepo: fakeTwitterRepo}
}

func (pts FollowUserService) FollowUserCaller(user, followedUser int) error {
	return pts.fakeTwitterRepo.FollowUser(user, followedUser)
}
