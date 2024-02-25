package handlers

import (
	"io"
	"net/http"

	"github.com/alexandr-orlov/go-short/config"
	"github.com/alexandr-orlov/go-short/internal/urldb"
	"github.com/go-chi/chi"
)

var udb = make(urldb.Urldb)

func GetHandler(res http.ResponseWriter, req *http.Request) {

	// do get log
	id := chi.URLParam(req, "id")

	url, err := udb.Get(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func GetRootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusBadRequest)
}

func PostHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "BadRequest", http.StatusBadRequest)
		}

		id, err := udb.Create(string(body))
		if err != nil {
			http.Error(res, "BadRequest", http.StatusBadRequest)
		}

		// status code 201
		res.WriteHeader(http.StatusCreated)
		url := config.BaseURL + id
		res.Write([]byte(url))
		return
	} else {
		// status code 400
		res.WriteHeader(http.StatusBadRequest)
		return
	}
}
