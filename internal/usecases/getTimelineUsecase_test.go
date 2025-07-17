package usecases

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/juaniferro/fake-twitter/internal/models"
)

type mockGetTimelineService struct {
    response []models.TimelineTweet
    err      error
}

func (m mockGetTimelineService) GetTimelineCaller(user int) (tweets []models.TimelineTweet, err error) {
    return m.response, m.err
}

func TestTimelineGetter(t *testing.T) {
    tests := []struct {
        name     string
        fields   GetTimelineServiceInterface
        args     int
        want     []models.TimelineTweet
        wantErr  bool
    }{
        {
            name: "success getting timeline from service",
            fields: mockGetTimelineService{
                response: []models.TimelineTweet{
                    {Username: "1", Content: "Hello", CreatedAt: time.Date(2025, 7, 16, 0, 0, 0, 0, time.UTC)},
                },
                err: nil,
            },
            args:    1,
            want:    []models.TimelineTweet{{Username: "1", Content: "Hello", CreatedAt: time.Date(2025, 7, 16, 0, 0, 0, 0, time.UTC)}},
            wantErr: false,
        },
        {
            name: "error from service",
            fields: mockGetTimelineService{
                response: nil,
                err:      errors.New("db error"),
            },
            args:    2,
            want:    nil,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gtu := GetTimelineUsecase{
                GetTimelineService: tt.fields,
            }
            got, err := gtu.TimelineGetter(tt.args)
            if (err != nil) != tt.wantErr {
                t.Errorf("TimelineGetter() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
                t.Errorf("TimelineGetter() got = %v, want %v", got, tt.want)
            }
        })
    }
}