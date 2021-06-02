package glab

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"pingmen/config"
	"strconv"
	"strings"

	"github.com/xanzy/go-gitlab"
)

// Init - initialize webhook
func Init(cfg *config.Config) *Webhook {
	log.Printf("Glab:Init: start")
	defer log.Printf("Glab:Init: inited")

	w := Webhook{
		// TODO: задел под дополнительные события
		Events: []gitlab.EventType{gitlab.EventTypeMergeRequest},
		Config: cfg,
	}

	return &w
}

// Run - run webhook
func (w *Webhook) Run() {
	log.Printf("Glab:Run: start")
	defer log.Printf("Glab:Run: end")

	serveMux := http.NewServeMux()
	serveMux.Handle(w.Config.Gitlab.WebhookMethod, w)

	if err := http.ListenAndServe(w.listenPath(), serveMux); err != nil {
		log.Fatalf("Glab:Run: HTTP server error: %v", err)
	}
}

// listenPath - create path for ListenAndServe
func (w *Webhook) listenPath() string {
	var build strings.Builder

	build.WriteString(w.Config.Gitlab.WebhookHost)
	build.WriteString(":")
	build.WriteString(strconv.Itoa(w.Config.Gitlab.WebhookPort))

	return build.String()
}

func (w *Webhook) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	event, err := w.parse(request)
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(fmt.Sprintf("could parse the webhook event: %v", err)))
		return
	}

	//	TODO

}

// parse - parse webhooks
func (w *Webhook) parse(r *http.Request) (interface{}, error) {
	defer func() {
		if _, err := io.Copy(ioutil.Discard, r.Body); err != nil {
			log.Printf("could discard request body: %v", err)
		}
		if err := r.Body.Close(); err != nil {
			log.Printf("could not close request body: %v", err)
		}
	}()

	if r.Method != http.MethodPost {
		return nil, errors.New("invalid HTTP Method")
	}

	if r.Header.Get(xGitlabToken) != w.Config.Gitlab.Token {
		return nil, errors.New("token validation failed")
	}

	if strings.TrimSpace(r.Header.Get(xGitlabEvent)) == "" {
		return nil, errors.New("missing X-Gitlab-Event Header")
	}

	if gitlab.EventType(r.Header.Get(xGitlabEvent)) != w.Events[0] {
		return nil, errors.New("event not defined to be parsed")
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, errors.New("error reading request body")
	}

	return gitlab.ParseWebhook(gitlab.EventType(r.Header.Get(xGitlabEvent)), payload)
}
