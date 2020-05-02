package main

import "regexp"

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
