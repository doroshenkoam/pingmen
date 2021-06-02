package main

import (
	"flag"
	"log"
	"os"

	"pingmen/config"
)

func main() {

	var (
		exitCode = 1
		cfg      config.Config
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
		log.Printf("Log file is: %s", *logFile)

		lf, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
		if err != nil {
			log.Printf("Error opening logfile: %s", err)
			_, usage := flag.UnquoteUsage(flg.Lookup("o"))
			log.Printf("Usage: %v", usage)
			os.Exit(exitCode)
		}
		defer lf.Close()

		log.SetOutput(lf)
	}

	if *helpFlag {
		flg.PrintDefaults()
		os.Exit(exitCode)
	}
	exitCode++

	if err := config.Load(cfgFile, &cfg); err != nil {
		log.Printf("Config file unmarshal error: %s", err)
		_, usage := flag.UnquoteUsage(flg.Lookup("c"))
		log.Printf("Usage: %v", usage)
		os.Exit(exitCode)
	}
	exitCode++

}
