package main

import (
	"flag"

	"github.com/alexandr-orlov/go-short/config"
)

func parseFlags() {
	flag.StringVar(&config.ListenAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&config.BaseURL, "b", "http://localhost:8080/", "Base URL on retrun Location header")

	flag.Parse()
}
