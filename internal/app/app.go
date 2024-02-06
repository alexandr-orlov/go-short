package app

import (
	"net/http"

	"github.com/alexandr-orlov/go-short/internal/handlers"
	"github.com/alexandr-orlov/go-short/internal/urldb"
)

func Run() {

	udb := make(urldb.Urldb)

	http.HandleFunc("/", handlers.MakeRootHandler(udb))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
