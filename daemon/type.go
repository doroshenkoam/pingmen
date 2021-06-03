package daemon

import (
	"pingmen/config"

	"github.com/xanzy/go-gitlab"
)

// Typ - daemons type
type Typ struct {
	cfg         *config.Config
	mrToBotChan <-chan *gitlab.MergeEvent
	doneChan    <-chan struct{}
}
