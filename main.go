package main

import (
	"AISale/api"
	"AISale/config"
)

func main() {
	//settings, err := config.LoadENV()
	//if err != nil {
	//	log.Fatal(err)
	//}

	api.RouterStart(config.Settings{})
}
