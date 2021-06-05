package glab

import (
	"pingmen/config"
	"sync"

	"github.com/xanzy/go-gitlab"
)

const (
	xGitlabEvent = "X-Gitlab-Event"
	xGitlabToken = "X-Gitlab-Token"
)

// Webhook - base webhook struct (exclude events)
type Webhook struct {
	event       gitlab.EventType
	config      *config.Config
	mrToBotChan chan *gitlab.MergeEvent
	doneChan    <-chan struct{}
	wg          *sync.WaitGroup
}
