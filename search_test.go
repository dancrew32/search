package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/search", nil)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("q", "hello world")
	req.URL.RawQuery = params.Encode()
	assert.Equal(t, req.URL.String(), "/search?q=hello+world")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchHandler)

	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)

	output := recorder.Body.String()
	expected := "Gorilla!\n"
	assert.Equal(t, output, expected)
}
