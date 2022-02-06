package main

import (
	"os"

	"github.com/rs/zerolog/log"

	// "golang.org/x/oauth2"
	// "golang.org/x/oauth2/clientcredentials"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	// oauth2 configures a client that uses app credentials to keep a fresh token
	// config := &clientcredentials.Config{
	// 	ClientID:     os.Getenv("TWITTER_CONSUMER_KEY"),
	// 	ClientSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
	// 	TokenURL:     "https://api.twitter.com/oauth2/token",
	// }
	// http.Client will automatically authorize Requests
	// httpClient := config.Client(oauth2.NoContext)

	config := oauth1.NewConfig(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	c := twitter.NewClient(httpClient)

	_, _, err := c.Statuses.Retweet(1365504377885241348, &twitter.StatusRetweetParams{})
	if err != nil {
		log.Error().Err(err).Msg("error retweeting")
	}
}
