package routers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alexandr-orlov/go-short/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testRequest(t *testing.T, ts *httptest.Server, method,
	path, data string) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, strings.NewReader(data))
	require.NoError(t, err)

	client := ts.Client()

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func TestRouter(t *testing.T) {
	config.ParseFlags()
	ts := httptest.NewServer(URLRouter())
	defer ts.Close()

	var testTable = []struct {
		url            string
		want           string
		method         string
		status         int
		data           string
		locationHeader string
	}{
		{"/", "", http.MethodGet, http.StatusBadRequest, "", ""},
		{"/", "http://localhost:8080/e98192e1", http.MethodPost, http.StatusCreated, "https://ya.ru", ""},
		{"/e98192e1", "", http.MethodGet, http.StatusTemporaryRedirect, "", "https://ya.ru"},
		{"/someid1", "", http.MethodGet, http.StatusBadRequest, "", ""},
	}
	for _, v := range testTable {
		resp, get := testRequest(t, ts, v.method, v.url, v.data)
		defer resp.Body.Close()

		assert.Equal(t, v.status, resp.StatusCode)
		assert.Equal(t, v.want, get)
		assert.Equal(t, v.locationHeader, resp.Header.Get("Location"))
	}
}
