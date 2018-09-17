package main

import (
	"github.com/IT-Accelerator-Tweet/src/domain"
	"github.com/IT-Accelerator-Tweet/src/service"
	"github.com/abiosoft/ishell"
	"strconv"
)

func main() {

	tm := service.NewTweetManager()

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type your username: ")

			user := c.ReadLine()

			c.Print("Type your tweet: ")

			text := c.ReadLine()

			tweet := domain.NewTextTweet(user, text)

			id, err := tm.PublishTweet(tweet)

			if err == nil {
				c.Printf("Tweet sent with id: %v\n", id)
			} else {
				c.Print("Error publishing tweet:", err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows the last tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := tm.GetTweet()

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all the tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tm.GetTweets()

			for _, tweet := range tweets {
				c.Println(tweet)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetById",
		Help: "Shows the tweet with the provided id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type the id: ")

			id, _ := strconv.Atoi(c.ReadLine())

			tweet := tm.GetTweetById(id)

			c.Println(tweet)

			return
		},
	})

	shell.Run()
}