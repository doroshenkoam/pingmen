package daemon

import "log"

// Receiver - analysis merge requests and sending to the bot
func (t *Typ) Receiver() {
	log.Printf("Glab:Receiver: daemon Receiver start")
	defer log.Printf("Glab:Receiver: daemon receiver end")

	for {
		select {
		case mr := <-t.mrToBotChan:
			if !t.isExistedProject(mr.Project.Name) {
				continue
			}

			if !t.isAction(mr.ObjectAttributes.Action) {
				continue
			}

			// TODO:
			//mr.ObjectAttributes.URL
			//mr.ObjectAttributes.Title
			//mr.ObjectAttributes.Description

		case <-t.doneChan:
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
