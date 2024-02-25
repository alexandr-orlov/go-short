package app

import (
	"fmt"
	"net/http"

	"github.com/alexandr-orlov/go-short/config"
	"github.com/alexandr-orlov/go-short/internal/routers"
)

func Run() {
	fmt.Println("Running server on: ", config.ListenAddr)
	http.ListenAndServe(config.ListenAddr, routers.URLRouter())

}
