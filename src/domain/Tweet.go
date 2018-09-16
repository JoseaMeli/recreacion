package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date time.Time
}

func NewTweet(user string, text string)*Tweet{
	t := time.Now()
	tweet :=  Tweet{user, text, t}
	return &tweet
}