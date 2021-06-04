package daemon

import (
	"pingmen/config"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/xanzy/go-gitlab"
)

// Typ - daemons type
type Typ struct {
	cfg         *config.Config
	bot         *tgbotapi.BotAPI
	wg          *sync.WaitGroup
	mrToBotChan <-chan *gitlab.MergeEvent
	doneChan    <-chan struct{}
}
