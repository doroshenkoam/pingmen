package daemon

import (
	"log"
	"pingmen/config"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/xanzy/go-gitlab"
)

// Init - init daemons
func Init(cfg *config.Config, bot *tgbotapi.BotAPI, wg *sync.WaitGroup,
	mrToBotChan <-chan *gitlab.MergeEvent, doneChan <-chan struct{}) *Typ {
	log.Printf("Daemon:Init: start")
	defer log.Printf("Daemon:Init: inited")

	return &Typ{
		cfg:         cfg,
		bot:         bot,
		wg:          wg,
		mrToBotChan: mrToBotChan,
		doneChan:    doneChan,
	}
}
