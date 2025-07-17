package handlers

import (
    "bytes"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"
)

type mockPostTweetUsecase struct {
    err error
}

func (m mockPostTweetUsecase) TweetPoster(user int, tweet string) error {
    return m.err
}

func TestHandlePostTweet(t *testing.T) {
    tests := []struct {
        name           string
        userIDHeader   string
        body           string
        usecase        PostTweetUsecaseInterface
        wantStatusCode int
    }{
        {
            name:           "success",
            userIDHeader:   "1",
            body:           `{"content":"Hello world"}`,
            usecase:        mockPostTweetUsecase{err: nil},
            wantStatusCode: http.StatusCreated,
        },
        {
            name:           "invalid user_id header",
            userIDHeader:   "abc",
            body:           `{"content":"Hello world"}`,
            usecase:        mockPostTweetUsecase{err: nil},
            wantStatusCode: http.StatusBadRequest,
        },
        {
            name:           "invalid json body",
            userIDHeader:   "1",
            body:           `not a json`,
            usecase:        mockPostTweetUsecase{err: nil},
            wantStatusCode: http.StatusBadRequest,
        },
        {
            name:           "usecase returns error",
            userIDHeader:   "1",
            body:           `{"content":"Hello world"}`,
            usecase:        mockPostTweetUsecase{err: errors.New("tweet can´t be longer than 280 characters")},
            wantStatusCode: http.StatusBadRequest,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            handler := PostTweetHandler{postTweetUsecase: tt.usecase}
            req := httptest.NewRequest("POST", "/tweet", bytes.NewBufferString(tt.body))
            req.Header.Set("user_id", tt.userIDHeader)
            req.Header.Set("Content-Type", "application/json")
            rr := httptest.NewRecorder()

            handler.HandlePostTweet(rr, req)

            if rr.Code != tt.wantStatusCode {
                t.Errorf("status code = %v, want %v", rr.Code, tt.wantStatusCode)
            }
        })
    }
}