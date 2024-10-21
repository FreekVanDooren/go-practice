package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestKeyEndpointHandler(t *testing.T) {
	t.Run("GET endpoint on empty store", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(keyEndpointHandler(make(map[string]string)))

		handler.ServeHTTP(recorder, req)

		body, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatal(err)
		}
		gotCode := recorder.Code
		gotBody := string(body[:])
		if gotCode != 200 || gotBody != "Store is empty" {
			t.Errorf("Expected 200: Store is empty, got: %v: %v", gotCode, body)
		}
	})

	t.Run("POST endpoint", func(t *testing.T) {
		input := "{\"beast\": \"value\"}"
		req := httptest.NewRequest(http.MethodPost, "/key", strings.NewReader(input))
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(keyEndpointHandler(make(map[string]string)))

		handler.ServeHTTP(recorder, req)

		body, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatal(err)
		}
		gotCode := recorder.Code
		gotBody := string(body[:])
		if gotCode != 200 || gotBody != "map[beast:wild]" {
			t.Errorf("Expected 200: Store is empty, got: %v: %v", gotCode, body)
		}
	})
}
