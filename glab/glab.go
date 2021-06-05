package glab

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"pingmen/config"
	"strconv"
	"strings"
	"sync"

	"github.com/xanzy/go-gitlab"
)

// Init - initialize webhook
func Init(cfg *config.Config, mrChan chan *gitlab.MergeEvent, doneChan <-chan struct{}, wg *sync.WaitGroup) *Webhook {
	log.Printf("Glab:Init: start")
	defer log.Printf("Glab:Init: inited")

	w := Webhook{
		event:       gitlab.EventTypeMergeRequest,
		config:      cfg,
		mrToBotChan: mrChan,
		doneChan:    doneChan,
		wg:          wg,
	}

	return &w
}

// Run - run webhook
func (w *Webhook) Run() {
	log.Printf("Glab:Run: start")
	defer log.Printf("Glab:Run: end")

	server := &http.Server{Addr: w.listenPath(), Handler: w}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Glab:Run: HTTP server error: %v", err)
		}

	}()

	w.wg.Add(1)
	go func() {
		defer w.wg.Done()

		for {
			select {
			case <-w.doneChan:
				if err := server.Shutdown(context.Background()); err != nil {
					log.Fatalf("Glab:Run: server shutdown error: %v\nkill process manual", err)
					return
				}

				log.Printf("Glab:Run: webhook server: end")
			}
		}
	}()

}

// listenPath - create path for ListenAndServe
func (w *Webhook) listenPath() string {
	var build strings.Builder
	defer build.Reset()

	build.WriteString(":")
	build.WriteString(strconv.Itoa(w.config.Gitlab.WebhookPort))

	return build.String()
}

func (w *Webhook) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	mr, err := w.validate(request)
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(fmt.Sprintf("could parse the webhook event: %v", err)))
		return
	}

	// send mr to FP receiver
	w.mrToBotChan <- mr

	writer.WriteHeader(204)
}

// validate - validate webhooks
func (w *Webhook) validate(r *http.Request) (*gitlab.MergeEvent, error) {
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

	if r.Header.Get(xGitlabToken) != w.config.Gitlab.Token {
		return nil, errors.New("token validation failed")
	}

	if strings.TrimSpace(r.Header.Get(xGitlabEvent)) == "" {
		return nil, errors.New("missing X-Gitlab-Event Header")
	}

	if gitlab.EventType(r.Header.Get(xGitlabEvent)) != w.event {
		return nil, errors.New("event not defined to be parsed")
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, errors.New("error reading request body")
	}

	return parseBody(payload)
}

// parseBody - parse body merge_request response
func parseBody(payload []byte) (*gitlab.MergeEvent, error) {
	var mr gitlab.MergeEvent

	if err := json.Unmarshal(payload, mr); err != nil {
		return nil, err
	}

	return &mr, nil
}
