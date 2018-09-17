package domain

import (
	"fmt"
	"time"
)

type QuoteTweet struct {
	Id   int
	User string
	Text string
	Date *time.Time
	QuotedTweet *TextTweet
}

func NewQuoteTweet(user string, text string, quotedTweet *TextTweet) *QuoteTweet {
	t := time.Now()
	tweet := QuoteTweet{ContadorDeTweets,user,text, &t, quotedTweet}

	var puntero *QuoteTweet
	puntero = &tweet

	ContadorDeTweets ++

	return puntero
}

func (t *QuoteTweet) PrintableTweet() string{
	return fmt.Sprintf("@%s: %s \"%s\"", t.User, t.Text, t.QuotedTweet.PrintableTweet())
}

func (t *QuoteTweet) String() string {
	return fmt.Sprintf("@%s: %s \"%s\"", t.User, t.Text, t.QuotedTweet.PrintableTweet())
}