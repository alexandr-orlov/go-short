package app

import (
	"net/http"

	"github.com/alexandr-orlov/go-short/internal/routers"
)

func Run() {

	http.ListenAndServe(":8080", routers.URLRouter())

}
