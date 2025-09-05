package main

import (
	"AISale/api"
	"AISale/config"
	"log"
)

func main() {
	settings, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	api.RouterStart(settings)
}
