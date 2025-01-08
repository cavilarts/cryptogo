package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from go"))
	})
	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Println("The server couldn't start")
	}
}