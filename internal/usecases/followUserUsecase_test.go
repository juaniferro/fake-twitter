package usecases

import (
    "errors"
    "testing"
)

type mockFollowUserService struct {
    err error
}

func (m mockFollowUserService) FollowUserCaller(user, followedUser int) error {
    return m.err
}

func TestUserFollower(t *testing.T) {
    tests := []struct {
        name    string
        service FollowUserServiceInterface
        user    int
        follow  int
        wantErr bool
    }{
        {
            name:    "success following user",
            service: mockFollowUserService{err: nil},
            user:    1,
            follow:  2,
            wantErr: false,
        },
        {
            name:    "service returns error",
            service: mockFollowUserService{err: errors.New("already following")},
            user:    1,
            follow:  2,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            usecase := FollowUserUsecase{followUserService: tt.service}
            err := usecase.UserFollower(tt.user, tt.follow)
            if (err != nil) != tt.wantErr {
                t.Errorf("UserFollower() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}