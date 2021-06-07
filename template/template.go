package template

import (
	"io/ioutil"
	"pingmen/config"
	"pingmen/logWrap"
	"strconv"
	"strings"

	"github.com/xanzy/go-gitlab"
)

// Load - loading template
func Load(path string) (string, error) {
	var (
		logger = logWrap.SetBaseFields("template", "Load")
	)

	logger.Info("Start")
	defer logger.Info("Inited")

	logger.Info("Config file is: ", path)

	templateData, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(templateData), nil
}

// CreateMsg - create message by template
func CreateMsg(cfg *config.Config, tmpl string, mr *gitlab.MergeEvent) string {
	var (
		logger = logWrap.SetBaseFields("template", "CreateMsg")
	)

	logger.WithField("template", tmpl).Debug("Start")
	defer logger.WithField("template", tmpl).Debug("Inited")

	tmpl = strings.ReplaceAll(tmpl, "{{Project Name}}", mr.Project.Name)
	tmpl = strings.ReplaceAll(tmpl, "{{Project Description}}", mr.Project.Description)
	tmpl = strings.ReplaceAll(tmpl, "{{ObjectAttributes Action}}", mr.ObjectAttributes.Action)
	tmpl = strings.ReplaceAll(tmpl, "{{ObjectAttributes AuthorID}}", strconv.Itoa(mr.ObjectAttributes.AuthorID))
	tmpl = strings.ReplaceAll(tmpl, "{{ObjectAttributes MergeUserID}}", strconv.Itoa(mr.ObjectAttributes.MergeUserID))
	tmpl = strings.ReplaceAll(tmpl, "{{ObjectAttributes MergeError}}", mr.ObjectAttributes.MergeError)
	tmpl = strings.ReplaceAll(tmpl, "{{ObjectAttributes MergeStatus}}", mr.ObjectAttributes.MergeStatus)
	tmpl = strings.ReplaceAll(tmpl, "{{ObjectAttributes Title}}", mr.ObjectAttributes.Title)
	tmpl = strings.ReplaceAll(tmpl, "{{ObjectAttributes URL}}", mr.ObjectAttributes.URL)
	tmpl = strings.ReplaceAll(tmpl, "{{ObjectAttributes Description}}", mr.ObjectAttributes.Description)
	tmpl = strings.ReplaceAll(tmpl, "{{Users}}", cfg.Users.Field)

	return tmpl
}
