package handlers

import (
	"io"
	"net/http"

	"github.com/alexandr-orlov/go-short/internal/urldb"
)

func MakeRootHandler(udb urldb.Urldb) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			PostHandler(res, req, udb)
		} else if req.Method == http.MethodGet {
			GetHandler(res, req, udb)
		} else {
			http.Error(res, "BadRequest", http.StatusBadRequest)
		}
	}
}

func GetHandler(res http.ResponseWriter, req *http.Request, udb urldb.Urldb) {
	if req.URL.Path == "/" {
		res.WriteHeader(http.StatusBadRequest)
		return
	} else {
		// do get log
		id := req.URL.Path[1:]
		url, err := udb.Get(id)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		res.Header().Set("Location", url)
		res.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
}

func PostHandler(res http.ResponseWriter, req *http.Request, udb urldb.Urldb) {
	if req.URL.Path == "/" {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "BadRequest", http.StatusBadRequest)
		}

		id, err := udb.Add(string(body))
		if err != nil {
			http.Error(res, "BadRequest", http.StatusBadRequest)
		}

		// status code 201
		res.WriteHeader(http.StatusCreated)
		url := "http://localhost:8080/" + id
		res.Write([]byte(url))
		return
	} else {
		// status code 400
		res.WriteHeader(http.StatusBadRequest)
		return
	}
}
