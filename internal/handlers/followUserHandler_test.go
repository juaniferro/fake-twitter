package handlers

import (
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
)

type mockFollowUserUsecase struct {
    err error
}

func (m mockFollowUserUsecase) UserFollower(user, followedUser int) error {
    return m.err
}

func TestHandleFollowUser(t *testing.T) {
    tests := []struct {
        name           string
        userIDHeader   string
        pathParam      string
        usecase        FollowUserUsecaseInterface
        wantStatusCode int
        wantBody       string
    }{
        {
            name:           "success follow",
            userIDHeader:   "1",
            pathParam:      "2",
            usecase:        mockFollowUserUsecase{err: nil},
            wantStatusCode: http.StatusCreated,
            wantBody:       "",
        },
        {
            name:           "invalid user_id header",
            userIDHeader:   "abc",
            pathParam:      "2",
            usecase:        mockFollowUserUsecase{err: nil},
            wantStatusCode: http.StatusBadRequest,
            wantBody:       "Invalid user_id\n",
        },
        {
            name:           "invalid followed user id",
            userIDHeader:   "1",
            pathParam:      "xyz",
            usecase:        mockFollowUserUsecase{err: nil},
            wantStatusCode: http.StatusBadRequest,
            wantBody:       "Invalid followed user id\n",
        },
        {
            name:           "usecase returns error",
            userIDHeader:   "1",
            pathParam:      "2",
            usecase:        mockFollowUserUsecase{err: errors.New("already following")},
            wantStatusCode: http.StatusBadRequest,
            wantBody:       "already following\n",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            handler := FollowUserHandler{followUserUsecase: tt.usecase}
            req := httptest.NewRequest("POST", "/follow/"+tt.pathParam, nil)
            req = mux.SetURLVars(req, map[string]string{"followed_user_id": tt.pathParam})
            req.Header.Set("user_id", tt.userIDHeader)
            rr := httptest.NewRecorder()

            handler.HandleFollowUser(rr, req)

            if rr.Code != tt.wantStatusCode {
                t.Errorf("status code = %v, want %v", rr.Code, tt.wantStatusCode)
            }
            if rr.Body.String() != tt.wantBody {
                t.Errorf("body = %q, want %q", rr.Body.String(), tt.wantBody)
            }
        })
    }
}