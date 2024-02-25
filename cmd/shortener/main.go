package main

import (
	"github.com/alexandr-orlov/go-short/config"
	"github.com/alexandr-orlov/go-short/internal/app"
)

func main() {
	config.ParseFlags()
	app.Run()
}
