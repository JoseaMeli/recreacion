package service

import (
	"errors"
	"github.com/recreacion/src/domain"

		)

var Tweet *domain.Tweet

func PublishTweet(tweet *domain.Tweet) error{
	if tweet.User == "" {
		return errors.New("user is required")
	} else if tweet.Text == "" {
		return errors.New("text is required")
	} else if len(tweet.Text) > 140{
		return errors.New("text exceeds 140 characters")
	} else {
		Tweet = tweet
		return nil
	}
}

func GetTweet() *domain.Tweet{
	return Tweet
}