package daemon

import (
	"fmt"
	"pingmen/logWrap"
	"pingmen/template"
	"strings"

	"github.com/sirupsen/logrus"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/xanzy/go-gitlab"
)

// Receiver - analysis merge requests and sending to the chat
func (t *Typ) Receiver() {
	var (
		logger = logWrap.SetBaseFields("daemon", "Receiver")
	)

	logger.Info("Receiver start")
	defer logger.Info("Receiver end")

	for n := 1; n <= t.cfg.Telegram.WorkersCount; n++ {
		t.wg.Add(1)
		go t.receiverWorker(n)
	}
}

// receiverWorker - worker for receiver daemon
func (t *Typ) receiverWorker(n int) {
	var (
		logger = logWrap.SetBaseFields("daemon", fmt.Sprintf("receiverWorker-%d", n))
	)

	logger.Info("Start")

	defer func() {
		if err := recover(); err != nil {
			logger.Panic("PANIC: ", err)
		}
	}()

	for {
		select {
		case mr := <-t.mrToBotChan:
			logger.Debug("Get mr")

			if !t.isExistedProject(mr.Project.Name) {
				logger.WithField(
					"project_name", mr.Project.Name,
				).Debug("Skip mr")
				continue
			}

			if !t.isAction(mr.ObjectAttributes.Action) {
				logger.WithField(
					"action", mr.ObjectAttributes.Action,
				).Debug("Skip mr")
				continue
			}

			if t.templ != nil && *t.templ != "" {
				t.sendMsg(template.CreateMsg(t.cfg, *t.templ, mr))
				continue
			}

			t.sendMsg(t.createMsg(mr))

		case <-t.doneChan:
			logger.Info("End")
			t.wg.Done()
			return
		}
	}
}

// isExistedProject - check project name
func (t *Typ) isExistedProject(project string) bool {
	for i := range t.cfg.Projects.Dictionary {
		if project == t.cfg.Projects.Dictionary[i] {
			return true
		}
	}

	return false
}

// isAction - check merge action
func (t *Typ) isAction(action string) bool {
	for i := range t.cfg.Gitlab.Actions {
		if action == t.cfg.Gitlab.Actions[i] {
			return true
		}
	}

	return false
}

// createMsg - creating message about merge request operation
func (t *Typ) createMsg(mr *gitlab.MergeEvent) string {
	var msg strings.Builder
	defer msg.Reset()

	msg.WriteString(mr.ObjectAttributes.Action)
	msg.WriteString(": ")
	msg.WriteString(mr.ObjectAttributes.Title)
	msg.WriteString("\n")
	msg.WriteString(mr.ObjectAttributes.URL)
	msg.WriteString("\n")
	msg.WriteString(mr.ObjectAttributes.Description)
	msg.WriteString("\n")
	msg.WriteString(t.cfg.Users.Field)

	return msg.String()
}

// sendMsg - sending message to chat
func (t *Typ) sendMsg(msg string) {
	var (
		logger = logWrap.SetBaseFields("daemon", "sendMsg")
	)

	logger.WithFields(logrus.Fields{
		"chat_id": t.cfg.Telegram.ChatID,
		"message": msg,
	}).Debug("Send message")

	_, err := t.bot.Send(tg.NewMessage(t.cfg.Telegram.ChatID, msg))
	if err != nil {
		logger.WithField(
			"error", err,
		).Error("Send message")
	}
}
