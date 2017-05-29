package main

import (
	"time"
)

// message represets a single message
type message struct {
	Name    string
	Message string
	When    time.Time
}
