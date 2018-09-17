package service

import (
	"errors"
	"github.com/IT-Accelerator-Tweet/src/domain"
)

type TweetManager struct {
	TextTweet *domain.TextTweet
	ContenedorDeTweets []*domain.TextTweet
	Um *UsuarioManager
}

//-----------FUNCION--------------//

func NewTweetManager() *TweetManager{
	tm := TweetManager{}
	tm.ContenedorDeTweets = make([]*domain.TextTweet,0)
	tm.Um = NewUsuarioManager()

	tm.Um.AgregarUsuario(&domain.Usuario{"kei", "kei@meli.com", "kei", "123"})
	tm.Um.AgregarUsuario(&domain.Usuario{"jose", "jose@meli.com", "jose", "1234"})
	return &tm
}

//----------METODOS--------------//

func (tm *TweetManager)PublishTweet(tweet *domain.TextTweet) (int, error) {

	usuarioEncontrado := tm.Um.GetUsuarioByNombre(tweet.User)

	if usuarioEncontrado != nil {
		if tweet.User == "" {
			return 0, errors.New("user is required")
		} else if tweet.Text == "" {
			return 0, errors.New("text is required")
		} else if len(tweet.Text) > 140 {
			return 0, errors.New("text exceeds 140 characters")
		} else {
			tm.TextTweet = tweet
			tm.ContenedorDeTweets = append(tm.ContenedorDeTweets, tweet)
			return tweet.Id, nil
		}
	} else {
		return 0, errors.New("No existe el Usuario")
	}
}

func (tm *TweetManager)GetTweet() *domain.TextTweet {
	return tm.TextTweet
}

func (tm *TweetManager)GetTweets() []*domain.TextTweet {
	return tm.ContenedorDeTweets
}

func (tm *TweetManager)GetTweetById(id int)  *domain.TextTweet{
	for _,tweet := range tm.ContenedorDeTweets {
		if tweet.Id == id {
			return tweet
		}
	}
	return nil
}

func (tm *TweetManager)CountTweetsByUser(user string) int {
	var i int
	for _,tweet := range tm.ContenedorDeTweets {
		if tweet.User == user {
			i ++
		}
	}
	return i
}

func (tm *TweetManager)GetTweetsByUser(user string) []*domain.TextTweet {
	tweets := make([]*domain.TextTweet,0)
	for _,tweet := range tm.ContenedorDeTweets {
		if tweet.User == user {
			tweets = append(tweets, tweet)
		}
	}
	return tweets
}