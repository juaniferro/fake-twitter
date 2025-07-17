package services

import (
    "errors"
    "testing"
)

type mockFollowUserRepo struct {
    err error
}

func (m mockFollowUserRepo) FollowUser(user, followedUser int) error {
    return m.err
}

func TestFollowUserCaller(t *testing.T) {
    tests := []struct {
        name    string
        repo    FollowUserRepositoryInterface
        user    int
        follow  int
        wantErr bool
    }{
        {
            name:    "success following user",
            repo:    mockFollowUserRepo{err: nil},
            user:    1,
            follow:  2,
            wantErr: false,
        },
        {
            name:    "repo returns error",
            repo:    mockFollowUserRepo{err: errors.New("already following")},
            user:    1,
            follow:  2,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            svc := FollowUserService{fakeTwitterRepo: tt.repo}
            err := svc.FollowUserCaller(tt.user, tt.follow)
            if (err != nil) != tt.wantErr {
                t.Errorf("FollowUserCaller() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}