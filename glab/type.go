package glab

import (
	"pingmen/config"

	"github.com/xanzy/go-gitlab"
)

const (
	xGitlabEvent = "X-Gitlab-Event"
	xGitlabToken = "X-Gitlab-Token"
)

// Webhook - base webhook struct (exclude events)
type Webhook struct {
	Events []gitlab.EventType
	Config *config.Config
}
