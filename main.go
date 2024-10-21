package main

import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	name := query.Get("name")
	if name == "" {
		fmt.Fprintf(writer, "Hello!")
	} else {
		fmt.Fprintf(writer, "Hello %v!", name)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
