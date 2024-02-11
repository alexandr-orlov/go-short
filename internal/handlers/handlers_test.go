package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alexandr-orlov/go-short/internal/urldb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHandler(t *testing.T) {
	type want struct {
		code           int
		response       string
		contentType    string
		LocationHeader string
	}

	udb := make(urldb.Urldb)

	testCases := []struct {
		name    string
		request string
		want    want
	}{
		{
			name:    "GET / - BadRequest",
			request: "/",
			want: want{
				code:           http.StatusBadRequest,
				response:       "",
				contentType:    "",
				LocationHeader: "",
			},
		},
		{
			name:    "GET /e98192e1",
			request: "/e98192e1",
			want: want{
				code:           http.StatusTemporaryRedirect,
				response:       "",
				contentType:    "",
				LocationHeader: "https://ya.ru",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// Добавляем в базу /id
			_, err := udb.Create(test.want.LocationHeader)
			require.NoError(t, err)

			request := httptest.NewRequest(http.MethodGet, test.request, nil)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			GetHandler(w, request, udb)
			res := w.Result()

			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			locationHeader := res.Header.Get("Location")

			require.NoError(t, err)
			assert.Equal(t, test.want.response, string(resBody))
			assert.Equal(t, test.want.LocationHeader, locationHeader)
		})
	}
}

func TestPostHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}

	udb := make(urldb.Urldb)

	testCases := []struct {
		name     string
		request  string
		data_url string
		want     want
	}{
		{
			name:     "POST any url",
			request:  "/anyid",
			data_url: "",
			want: want{
				code:        http.StatusBadRequest,
				response:    "",
				contentType: "",
			},
		},

		{
			name:     "POST https://ya.ru",
			request:  "/",
			data_url: "https://ya.ru",
			want: want{
				code:        http.StatusCreated,
				response:    "http://localhost:8080/e98192e1",
				contentType: "text/plain",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			request := httptest.NewRequest(http.MethodPost, test.request, strings.NewReader(test.data_url))
			// создаём новый Recorder
			w := httptest.NewRecorder()
			PostHandler(w, request, udb)
			res := w.Result()

			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			assert.Equal(t, test.want.response, string(resBody))
		})
	}
}
