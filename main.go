package main

import (
	"flag"
	"os"
	"os/signal"
	"pingmen/config"
	"pingmen/daemon"
	"pingmen/glab"
	"pingmen/logWrap"
	"sync"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {

	var (
		exitCode = 1
		cfg      config.Config
		logger   = logWrap.SetBaseFields("main", "main")
	)

	flg := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	cfgFile := flg.String("c", "", "-c <path to сonfig file>")
	flg.StringVar(cfgFile, "cfg", "", "--config <path to сonfig file>")
	logFile := flg.String("l", "", "-l <path to log file>")
	flg.StringVar(logFile, "log", "", "--log <path to log file>")
	helpFlag := flg.Bool("h", false, "help flag usage")
	flg.BoolVar(helpFlag, "help", false, "help flag usage")
	flg.Parse(os.Args[1:])

	if *logFile != "" {
		logger.Info("Log file is: %s", *logFile)

		lf, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
		if err != nil {
			logger.WithField(
				"error", err,
			).Error("Error opening logfile")

			_, usage := flag.UnquoteUsage(flg.Lookup("l"))

			logger.Fatal("Usage: %v", usage)
			os.Exit(exitCode)
		}
		defer lf.Close()

		logrus.SetOutput(lf)
	}

	if *helpFlag {
		flg.PrintDefaults()
		os.Exit(exitCode)
	}
	exitCode++

	if err := config.Load(*cfgFile, &cfg); err != nil {
		logger.WithField(
			"error", err,
		).Error("Config file unmarshal error")

		_, usage := flag.UnquoteUsage(flg.Lookup("c"))

		logger.Fatal("Usage: %v", usage)
		os.Exit(exitCode)
	}
	exitCode++

	logWrap.SetLogLevel(cfg.Loglevel)

	bot, err := tg.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		logger.WithField(
			"error", err,
		).Info("Bot initing error")

		os.Exit(exitCode)
	}
	exitCode++

	logger.WithField(
		"bot_user_name", bot.Self.UserName,
	).Info("Authorized on account")

	if cfg.Telegram.Debug {
		bot.Debug = true

		logger.Info("Bot debug enabled")
	}

	doneChan := make(chan struct{})
	mrChan := make(chan *gitlab.MergeEvent)

	interrupter := make(chan os.Signal, 1)
	signal.Notify(interrupter, os.Interrupt)

	wg := sync.WaitGroup{}

	d := daemon.Init(&cfg, bot, &wg, mrChan, doneChan)
	d.Receiver()

	g := glab.Init(&cfg, mrChan, doneChan, &wg)
	g.Run()

mLoop:
	for {
		select {
		case <-interrupter:
			close(doneChan)
			wg.Wait()
			break mLoop
		}
	}
}
