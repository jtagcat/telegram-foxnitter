package main

import (
	"net/url"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jtagcat/util/retry"
	"github.com/rs/zerolog/log"
	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	for _, e := range []string{"CLIENT_ID", "CLIENT_SECRET", "ACCESS_TOKEN", "ACCESS_SECRET", "TELEGRAM_KEY", "NITTER_DOMAIN"} {
		if os.Getenv(e) == "" {
			log.Fatal().Str("variable", e).Msg("Required envrionment variable must not be empty")
		}
	}
	tgIDStr := os.Getenv("TELEGRAM_ID")
	var tgID int64
	var err error
	if tgIDStr != "" {
		tgID, err = strconv.ParseInt(tgIDStr, 10, 64)
		if err != nil {
			log.Fatal().Err(err).Msg("TELEGRAM_ID must be a number")
		}
	}

	config := oauth1.NewConfig(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	tw := twitter.NewClient(httpClient)

	tg, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_KEY"))
	if err != nil {
		log.Fatal().Err(err).
			Msg("Failed to initialize bot; is TELEGRAM_KEY valid?")
	}
	log.Info().
		Str("user", tg.Self.UserName).
		Msg("Authorization successful.")
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60

	updates, err := tg.GetUpdatesChan(update)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed getting Telegram updates")
	}
	for u := range updates {
		if u.ChannelPost != nil { // message on channel
			u.Message = u.ChannelPost
		}
		if u.Message != nil { // update iz message with text
			if tgIDStr == "" {
				log.Info().Int64("id", u.Message.Chat.ID).Str("name", u.Message.Chat.Title).Msg("No ID set, relaying ID of incoming message")
			}
			if u.Message.Chat.ID == tgID && u.Message.Entities != nil {
				for _, e := range *u.Message.Entities {
					if e.Type == "url" {
						url, err := url.Parse(u.Message.Text[e.Offset : e.Offset+e.Length])
						if err != nil {
							log.Error().Err(err).Msg("Failed parsing URL")
							break
						}

						pathSlice := strings.Split(strings.Trim(url.Path, "/"), "/")
						if len(pathSlice) >= 3 && strings.HasPrefix(strings.ToLower(url.Host), strings.ToLower(os.Getenv("NITTER_DOMAIN"))) && string(pathSlice[1]) == "status" {
							twid, err := strconv.ParseInt(string(pathSlice[2]), 10, 64)
							if err != nil {
								log.Error().Err(err).Msg("Failed parsing tweet ID")
								break
							}

							err = retry.OnError(wait.Backoff{
								Duration: 2,
								Factor:   2,
								Steps:    4,
							}, func() (bool, error) {
								_, _, err := tw.Statuses.Retweet(twid, &twitter.StatusRetweetParams{})
								return true, err
							})
							if err != nil {
								log.Error().Err(err).Msg("error retweeting")
								break
							}
							log.Info().Int64("id", twid).Msg("retweeted")
						}
						break
					}
				}
			}
		}
	}
}
