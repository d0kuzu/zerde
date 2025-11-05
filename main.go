package main

import (
	"AISale/api"
	"AISale/cleanup"
	"AISale/config"
	"AISale/services/chrome"
	"AISale/services/jobs"
	"log"
)

func main() {
	settings, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	chromeClient := chrome.Init(settings.)

	cm := &cleanup.CleanupManager{}
	cm.Add(chromeClient.Close)
	go cm.Start()

	app := config.NewApp(chromeClient, settings)

	go jobs.CheckWaitingChats(app)

	api.RouterStart(app)
}
