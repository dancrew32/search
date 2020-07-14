package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/search", nil)
	assert.Nil(t, err)

	params := url.Values{"search": []string{"hello world"}}
	req.URL.RawQuery = params.Encode()
	assert.Equal(t, req.URL.String(), "/api/search?search=hello+world")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchHandler)

	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)

	var output SearchResponse
	err = json.NewDecoder(recorder.Body).Decode(&output)
	assert.Nil(t, err)
	assert.Equal(t, output.Data, "hello world")
}
