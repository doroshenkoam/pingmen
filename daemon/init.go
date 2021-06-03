package daemon

import (
	"pingmen/config"

	"github.com/xanzy/go-gitlab"
)

// Init - init daemons
func Init(cfg *config.Config, mrToBotChan <-chan *gitlab.MergeEvent, doneChan <-chan struct{}) *Typ {
	return &Typ{
		cfg:         cfg,
		mrToBotChan: mrToBotChan,
		doneChan:    doneChan,
	}
}
