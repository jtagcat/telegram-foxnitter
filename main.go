package main

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("TWITTER_CONSUMER_KEY"),
		ClientSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)
	// Twitter client
	c := twitter.NewClient(httpClient)

	_, _, err := c.Statuses.Retweet(1365504377885241348, &twitter.StatusRetweetParams{})
	if err != nil {
		log.Error().Err(err).Msg("error retweeting")
	}
}
