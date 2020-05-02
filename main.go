package main

import (
	"context"
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

	return bot, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load the environment variables")
	}

}
