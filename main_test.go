package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("basic", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(hello)

		handler.ServeHTTP(recorder, request)

		wanted := "Code: 200 and message: Hello!"
		got := fmt.Sprintf("Code: %v and message: %v", recorder.Code, recorder.Body)
		if got != wanted {
			t.Error(got)
		}
	})
	t.Run("with name", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		query := request.URL.Query()
		query.Add("name", "world")
		request.URL.RawQuery = query.Encode()
		handler := http.HandlerFunc(hello)

		handler.ServeHTTP(recorder, request)

		wanted := "Code: 200 and message: Hello world!"
		got := fmt.Sprintf("Code: %v and message: %v", recorder.Code, recorder.Body)
		if got != wanted {
			t.Error(got)
		}
	})
}
