package daemon

import (
	"log"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/xanzy/go-gitlab"
)

// Receiver - analysis merge requests and sending to the chat
func (t *Typ) Receiver() {
	log.Printf("Daemon:Receiver: daemon Receiver start")
	defer log.Printf("Daemon:Receiver: daemon receiver end")

	for n := 0; n >= t.cfg.Telegram.WorkersCount; n++ {
		t.wg.Add(1)
		go t.receiverWorker(n)
	}
}

// receiverWorker - worker for receiver daemon
func (t *Typ) receiverWorker(n int) {
	log.Printf("Daemon:receiverWorker: №%d start", n)

	defer func() {
		if err := recover(); err != nil {
			log.Printf("Daemon:receiverWorker: №%d PANIC: %#v", n, err)
		}
	}()

	for {
		select {
		case mr := <-t.mrToBotChan:
			if !t.isExistedProject(mr.Project.Name) {
				continue
			}

			if !t.isAction(mr.ObjectAttributes.Action) {
				continue
			}

			t.sendMsg(t.createMsg(mr))

		case <-t.doneChan:
			log.Printf("Daemon:receiverWorker: №%d end", n)
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

	msg.WriteString(mr.ObjectAttributes.Title)
	msg.WriteString("\n")
	msg.WriteString(mr.ObjectAttributes.URL)
	msg.WriteString("\n")
	msg.WriteString(mr.ObjectAttributes.Description)
	msg.WriteString("\n")

	for i := range t.cfg.Users.Dictionary {

		if i != 0 {
			msg.WriteString(" ")
		}

		msg.WriteString("@")
		msg.WriteString(t.cfg.Users.Dictionary[i])
	}

	return msg.String()
}

// sendMsg - sending message to chat
func (t *Typ) sendMsg(msg string) {
	log.Printf("Sending message (chat_id:%d): \n%s", t.cfg.Telegram.ChatID, msg)

	_, err := t.bot.Send(tg.NewMessage(t.cfg.Telegram.ChatID, msg))
	if err != nil {
		log.Fatalf("Sending message error: %s", err)
	}
}
