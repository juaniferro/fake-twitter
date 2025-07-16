package usecases

import (
	"errors"

	"github.com/juaniferro/fake-twitter/internal/services"
)

type PostTweetUsecase struct {
	postTweetService services.PostTweetService
}

type PostTweetUsecaseInterface interface{
	PostTweetCaller(user int, tweet string) error 
}

func NewPostTweetUsecase(postTweetService services.PostTweetService) *PostTweetUsecase {
	return &PostTweetUsecase{postTweetService: postTweetService}
}

func (ptu PostTweetUsecase) TweetPoster(user int, tweet string) error {
	if len(tweet) > 280 {
		return errors.New("tweet can´t be longer than 280 characters")
	}
	return ptu.postTweetService.PostTweetCaller(user, tweet)
}