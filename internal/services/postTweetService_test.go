package services

import (
    "errors"
    "testing"
)

type mockPostTweetRepo struct {
    err error
}

func (m mockPostTweetRepo) PostTweet(user int, tweet string) error {
    return m.err
}

func TestPostTweetCaller(t *testing.T) {
    tests := []struct {
        name    string
        repo    PostTweetRepositoryInterface
        user    int
        tweet   string
        wantErr bool
    }{
        {
            name:    "success posting tweet",
            repo:    mockPostTweetRepo{err: nil},
            user:    1,
            tweet:   "Hello world",
            wantErr: false,
        },
        {
            name:    "repo returns error",
            repo:    mockPostTweetRepo{err: errors.New("tweet too long")},
            user:    1,
            tweet:   "This tweet is way too long...",
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            svc := PostTweetService{fakeTwitterRepo: tt.repo}
            err := svc.PostTweetCaller(tt.user, tt.tweet)
            if (err != nil) != tt.wantErr {
                t.Errorf("PostTweetCaller() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}