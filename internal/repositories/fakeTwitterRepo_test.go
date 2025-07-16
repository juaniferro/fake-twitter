package repositories

import (
	"database/sql"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/juaniferro/fake-twitter/internal/models"
)

func TestFakeTwitterRepo_GetTimeline(t *testing.T) {
    type fields struct {
        setupMock func(mock sqlmock.Sqlmock)
    }
    tests := []struct {
        name      string
        fields    fields
        user      int
        want      []models.TimelineTweet
        wantErr   bool
    }{
        {
            name: "success getting timeline from db",
            fields: fields{
                setupMock: func(mock sqlmock.Sqlmock) {
                    rows := sqlmock.NewRows([]string{"name", "content", "created_at"}).
                        AddRow("Juani", "Hello", time.Date(2025, 7, 16, 0, 0, 0, 0, time.UTC))
                    mock.ExpectQuery(regexp.QuoteMeta(
                        "select users.name, tweets.content, tweets.created_at from tweets join users on tweets.tweet_user = users.id join user_follows on user_follows.followedUserID = users.id where user_follows.userID = ? order by tweets.created_at DESC",
                    )).WithArgs(1).WillReturnRows(rows)
                },
            },
            user: 1,
            want: []models.TimelineTweet{
                {Username: "Juani", Content: "Hello", CreatedAt: time.Date(2025, 7, 16, 0, 0, 0, 0, time.UTC)},
            },
            wantErr: false,
        },
        {
            name: "query error",
            fields: fields{
                setupMock: func(mock sqlmock.Sqlmock) {
                    mock.ExpectQuery(regexp.QuoteMeta(
                        "select users.name, tweets.content, tweets.created_at from tweets join users on tweets.tweet_user = users.id join user_follows on user_follows.followedUserID = users.id where user_follows.userID = ? order by tweets.created_at DESC",
                    )).WithArgs(2).WillReturnError(sql.ErrConnDone)
                },
            },
            user:    2,
            want:    nil,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            db, mock, err := sqlmock.New()
            if err != nil {
                t.Fatalf("failed to open sqlmock database: %v", err)
            }
            defer db.Close()

            tt.fields.setupMock(mock)
            repo := FakeTwitterRepo{db: db}
            got, err := repo.GetTimeline(tt.user)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetTimeline() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetTimeline() got = %v, want %v", got, tt.want)
            }
        })
    }
}