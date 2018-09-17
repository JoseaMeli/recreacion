package domain

import (
	"fmt"
	"time"
)

type ImageTweet struct {
	Id   int
	User string
	Text string
	Url string
	Date *time.Time
}

func NewImageTweet(user string, text string, url string) *ImageTweet {
	t := time.Now()
	tweet := ImageTweet{ContadorDeTweets,user,text, url, &t}

	var puntero *ImageTweet
	puntero = &tweet

	ContadorDeTweets ++

	return puntero
}

func (t *ImageTweet) PrintableTweet() string{
	return fmt.Sprintf("@%s: %s %s", t.User, t.Text,t.Url)
}

func (t *ImageTweet) String() string {
	return fmt.Sprintf("@%s: %s %s", t.User, t.Text,t.Url)
}
