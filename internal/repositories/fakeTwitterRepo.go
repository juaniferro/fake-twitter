package repositories

import (
	"database/sql"

	"github.com/juaniferro/fake-twitter/internal/models"
)

type FakeTwitterRepo struct {
	db *sql.DB
}

func NewFakeTwitterRepo(db *sql.DB) *FakeTwitterRepo{
	return &FakeTwitterRepo{db: db}
}

func (fk FakeTwitterRepo) PostTweet(user int, tweet string) error {
	 _, err := fk.db.Exec("INSERT INTO tweets (tweet_user, content) VALUES (?, ?)", user, tweet)
	 if err != nil{
		return err
	 }
	 return nil
}

func (fk FakeTwitterRepo) FollowUser(user, followedUser int) error {
	_, err := fk.db.Exec("INSERT INTO user_follows (userID, followedUserID) VALUES (?, ?)", user, followedUser)
	 if err != nil{
		return err
	 }
	 return nil
}

func (fk FakeTwitterRepo) GetTimeline(user int) (tweets []models.TimelineTweet, err error) {
	var timeline []models.TimelineTweet
	query := "select users.name, tweets.content, tweets.created_at from tweets join users on tweets.tweet_user = users.id join user_follows on user_follows.followedUserID = users.id where user_follows.userID = ? order by tweets.created_at DESC"
	rows, err := fk.db.Query(query, user)
    if err != nil {
        return nil, err
	}
    defer rows.Close()
  
    for rows.Next() {
        var tweet models.TimelineTweet
        if err := rows.Scan(&tweet.Username, &tweet.Content, &tweet.CreatedAt); err != nil {
            return nil, err
        }
        timeline = append(timeline, tweet)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return timeline, nil
}
