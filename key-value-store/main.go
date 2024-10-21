package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	keyValueStore := make(map[string]string)
	http.HandleFunc("/key", keyEndpointHandler(keyValueStore))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}

func keyEndpointHandler(keyValueStore map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			_, err := fmt.Fprintf(writer, fmt.Sprint(keyValueStore))
			if err != nil {
				fmt.Println(err)
			}
			return
		case http.MethodPost:
			body, err := io.ReadAll(request.Body)
			var data map[string]string
			err = json.Unmarshal(body, &data)
			if err != nil {
				// TODO: return 400
			}
			for key, value := range data {
				keyValueStore[key] = value
			}
			_, err = fmt.Fprintf(writer, fmt.Sprint(keyValueStore))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
