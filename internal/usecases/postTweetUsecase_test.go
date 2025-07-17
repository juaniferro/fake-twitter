package usecases

import (
    "errors"
    "testing"
)

type mockPostTweetService struct {
    err error
}

func (m mockPostTweetService) PostTweetCaller(user int, tweet string) error {
    return m.err
}

func TestTweetPoster(t *testing.T) {
    tests := []struct {
        name    string
        service PostTweetServiceInterface
        user    int
        tweet   string
        wantErr bool
    }{
        {
            name:    "success posting tweet",
            service: mockPostTweetService{err: nil},
            user:    1,
            tweet:   "Hello world",
            wantErr: false,
        },
        {
            name:    "service returns error",
            service: mockPostTweetService{err: errors.New("tweet can´t be longer than 280 characters")},
            user:    1,
            tweet:   "This tweet is way too long...",
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            usecase := PostTweetUsecase{postTweetService: tt.service}
            err := usecase.TweetPoster(tt.user, tt.tweet)
            if (err != nil) != tt.wantErr {
                t.Errorf("TweetPoster() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}