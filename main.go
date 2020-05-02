package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	ipfsCluster "github.com/ipfs/ipfs-cluster/api/rest/client"
	"github.com/joho/godotenv"
)

// Action that serves as a command for ipfs cluster
type Action string

var (
	actionRegexp        = regexp.MustCompile(`^\s*([[:graph:]]+)\s+(.+)`)
	pinRegexp           = regexp.MustCompile(`([[:graph:]]+)\s+([[:graph:]\s]+)`)
	PinAction    Action = "!pin"
	UnPinACtion  Action = "!unpin"
	AddAction    Action = "!add"
	HelpAction   Action = "!help"
)

// Bot is the twitter bot
type Bot struct {
	ctx           context.Context
	cancel        context.CancelFunc
	name          string
	id            string
	twClient      *twitter.Client
	clusterClient ipfsCluster.Client
	followedBy    sync.Map
	die           chan struct{}
}

// Kill destroys the bot and cancels the context
func (b Bot) Kill() {
	b.cancel()
}

// Name returns the bot's handle
func (b Bot) Name() string {
	return b.name
}

// ID returns the twitter user ID used by the bot
func (b Bot) ID() string {
	return b.id
}

func (b *Bot) fetchFollowers() {
	var nextCursor int64 = -1
	includeEntities := false

	for nextCursor != 0 {
		followers, _, err := b.twClient.Followers.List(
			&twitter.FollowerListParams{
				Count:               200,
				IncludeUserEntities: &includeEntities,
			},
		)

		if err != nil {
			log.Println(err)
		}

		fmt.Println(followers)

		nextCursor++
	}

}

// New creates Titanium Giant
func New() (*Bot, error) {
	ctx, cancel := context.WithCancel(context.Background())

	config := oauth1.NewConfig(os.Getenv("APIKEY"), os.Getenv("APISECRET"))
	token := oauth1.NewToken(os.Getenv("TOKEN"), os.Getenv("TOKENSECRET"))
	httpClient := config.Client(ctx, token)
	twClient := twitter.NewClient(httpClient)

	bot := &Bot{
		ctx:      ctx,
		cancel:   cancel,
		twClient: twClient,
		die:      make(chan struct{}, 1),
	}

	bot.fetchFollowers()
	// go bot.watchFollowers()
	// go bot.watchTweets()

	return bot, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load the environment variables")
	}

}
