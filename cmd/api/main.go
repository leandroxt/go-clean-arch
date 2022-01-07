package main

import (
	"flag"
	"log"

	"github.com/leandroxt/go-clean-arch/interface-adapters/inbound"
)

type config struct {
	giphyApiKey string
}

type application struct {
	handler inbound.Handler
}

func main() {
	var config config

	flag.StringVar(&config.giphyApiKey, "giphy-api-key", "-", "auth api integration")
	flag.Parse()

	app := &application{
		handler: inbound.NewHandler(config.giphyApiKey),
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
