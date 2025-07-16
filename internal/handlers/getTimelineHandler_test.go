package handlers

import (
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "time"

    "github.com/juaniferro/fake-twitter/internal/models"
)

type mockGetTimelineUsecase struct {
    timeline []models.TimelineTweet
    err      error
}

func (m mockGetTimelineUsecase) TimelineGetter(user int) ([]models.TimelineTweet, error) {
    return m.timeline, m.err
}

func TestHandleGetTimeline(t *testing.T) {
    tests := []struct {
        name           string
        userIDHeader   string
        usecase        GetTimelineUsecaseInterface
        wantStatusCode int
        wantBody       string
    }{
        {
            name:         "valid user, timeline returned",
            userIDHeader: "1",
            usecase: mockGetTimelineUsecase{
                timeline: []models.TimelineTweet{
                    {Username: "1", Content: "Hello", CreatedAt: time.Date(2025, 7, 16, 0, 0, 0, 0, time.UTC)},
                },
                err: nil,
            },
            wantStatusCode: http.StatusOK,
            wantBody:       `[{"Username":"1","Content":"Hello","CreatedAt":"2025-07-16T00:00:00Z"}]`,
        },
        {
            name:           "invalid user_id header",
            userIDHeader:   "abc",
            usecase:        mockGetTimelineUsecase{},
            wantStatusCode: http.StatusBadRequest,
            wantBody:       "Invalid user_id\n",
        },
        {
            name:         "usecase returns error",
            userIDHeader: "2",
            usecase: mockGetTimelineUsecase{
                timeline: nil,
                err:      errors.New("db error"),
            },
            wantStatusCode: http.StatusBadRequest,
            wantBody:       "db error\n",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            handler := GetTimelineHandler{getTimelineUsecase: tt.usecase}
            req := httptest.NewRequest("GET", "/timeline", nil)
            req.Header.Set("user_id", tt.userIDHeader)
            rr := httptest.NewRecorder()

            handler.HandleGetTimeline(rr, req)

            if rr.Code != tt.wantStatusCode {
                t.Errorf("status code = %v, want %v", rr.Code, tt.wantStatusCode)
            }

            gotBody := strings.TrimSpace(rr.Body.String())
            wantBody := strings.TrimSpace(tt.wantBody)

            if tt.wantStatusCode == http.StatusOK {
                var got, want []models.TimelineTweet
                if err := json.Unmarshal(rr.Body.Bytes(), &got); err != nil {
                    t.Fatalf("failed to unmarshal response: %v", err)
                }
                if err := json.Unmarshal([]byte(tt.wantBody), &want); err != nil {
                    t.Fatalf("failed to unmarshal wantBody: %v", err)
                }
                if len(got) != len(want) || got[0].Content != want[0].Content {
                    t.Errorf("response body = %v, want %v", got, want)
                }
            } else if gotBody != wantBody {
                t.Errorf("response body = %q, want %q", gotBody, wantBody)
            }
        })
    }
}