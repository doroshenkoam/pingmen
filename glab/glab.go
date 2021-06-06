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
	"pingmen/logWrap"
	"strconv"
	"strings"
	"sync"

	"github.com/xanzy/go-gitlab"
)

// Init - initialize webhook
func Init(cfg *config.Config, mrChan chan *gitlab.MergeEvent, doneChan <-chan struct{}, wg *sync.WaitGroup) *Webhook {
	var (
		logger = logWrap.SetBaseFields("glab", "Init")
	)

	logger.Info("Start")
	defer logger.Info("Inited")

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
	var (
		logger = logWrap.SetBaseFields("glab", "Init")
	)

	logger.Info("Start")
	defer logger.Info("End")

	server := &http.Server{Addr: w.listenPath(), Handler: w}

	go func() {
		logger.Info("Webhook start: listen port:%d", w.config.Gitlab.WebhookPort)

		if err := server.ListenAndServe(); err != nil {
			logger.WithField(
				"error", err,
			).Fatal("HTTP server error")
		}

	}()

	w.wg.Add(1)
	go func() {
		defer w.wg.Done()

		for {
			select {
			case <-w.doneChan:
				if err := server.Shutdown(context.Background()); err != nil {
					logger.WithField(
						"error", err,
					).Fatal("Server shutdown error\nkill process manual")
					return
				}

				logger.Info("End")
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
	var (
		logger = logWrap.SetBaseFields("glab", "ServeHTTP")
	)

	logger.Debug("Webhook used")

	mr, err := w.validate(request)
	if err != nil {
		log.Printf("Glab:ServeHTTP: error: %s", err)
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
	var (
		logger = logWrap.SetBaseFields("glab", "validate")
	)

	defer func() {
		if _, err := io.Copy(ioutil.Discard, r.Body); err != nil {
			logger.WithField(
				"error", err,
			).Error("Could discard request body")
		}
		if err := r.Body.Close(); err != nil {
			logger.WithField(
				"error", err,
			).Error("could not close request body")
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

	if err := json.Unmarshal(payload, &mr); err != nil {
		return nil, err
	}

	return &mr, nil
}
