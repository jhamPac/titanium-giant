package main

import (
	"context"
	"regexp"
	"sync"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/ipfs/ipfs-cluster/api/rest/client"
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
	clusterClient client.Client
	followedBy    sync.Map
	die           chan struct{}
}

// New creates Titanium Giant
func New() (*Bot, error) {
	return nil, nil
}

func main() {

}
