package services

import "github.com/juaniferro/fake-twitter/internal/repositories"

type PostTweetService struct {
	fakeTwitterRepo PostTweetRepositoryInterface
}

type PostTweetRepositoryInterface interface{
	PostTweet(user int, tweet string) error
}

func NewPostTweetService(fakeTwitterRepo repositories.FakeTwitterRepo) *PostTweetService{
	return &PostTweetService{fakeTwitterRepo: fakeTwitterRepo}
}

func (pts PostTweetService) PostTweetCaller(user int, tweet string) error {
	return pts.fakeTwitterRepo.PostTweet(user, tweet)
}