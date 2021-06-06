package daemon

import (
	"pingmen/config"
	"pingmen/logWrap"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/xanzy/go-gitlab"
)

// Init - init daemons
func Init(cfg *config.Config, bot *tgbotapi.BotAPI, wg *sync.WaitGroup,
	mrToBotChan <-chan *gitlab.MergeEvent, doneChan <-chan struct{}) *Typ {
	var (
		logger = logWrap.SetBaseFields("daemon", "Init")
	)

	logger.Info("Start")
	defer logger.Info("Inited")

	return &Typ{
		cfg:         cfg,
		bot:         bot,
		wg:          wg,
		mrToBotChan: mrToBotChan,
		doneChan:    doneChan,
	}
}
