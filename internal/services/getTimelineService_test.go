package services

import (
    "errors"
    "reflect"
    "testing"
    "time"

    "github.com/juaniferro/fake-twitter/internal/models"
)

// Mock for GetTimelineSRepositoryInterface
type mockTimelineRepo struct {
    tweets []models.TimelineTweet
    err    error
}

func (m mockTimelineRepo) GetTimeline(user int) ([]models.TimelineTweet, error) {
    return m.tweets, m.err
}

func TestGetTimelineCaller(t *testing.T) {
    tests := []struct {
        name      string
        repo      GetTimelineSRepositoryInterface
        user      int
        want      []models.TimelineTweet
        wantErr   bool
    }{
        {
            name: "success getting timeline from repo",
            repo: mockTimelineRepo{
                tweets: []models.TimelineTweet{
                    {Username: "1", Content: "Hello", CreatedAt: time.Date(2025, 7, 16, 0, 0, 0, 0, time.UTC)},
                },
                err: nil,
            },
            user:    1,
            want:    []models.TimelineTweet{{Username: "1", Content: "Hello", CreatedAt: time.Date(2025, 7, 16, 0, 0, 0, 0, time.UTC)}},
            wantErr: false,
        },
        {
            name: "repo error",
            repo: mockTimelineRepo{
                tweets: nil,
                err:    errors.New("db error"),
            },
            user:    2,
            want:    nil,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            svc := GetTimelineService{fakeTwitterRepo: tt.repo}
            got, err := svc.GetTimelineCaller(tt.user)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetTimelineCaller() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetTimelineCaller() got = %v, want %v", got, tt.want)
            }
        })
    }
}