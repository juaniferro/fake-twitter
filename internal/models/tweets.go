package models

import "time"

type Tweet struct {
	Id int
	TweetUser int
	Content string
	CreatedAt time.Time 
}

type TimelineTweet struct {
	Username string
	Content string
	CreatedAt time.Time 
}