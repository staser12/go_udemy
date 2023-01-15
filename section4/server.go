package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getgoing", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received!")
		fmt.Println(r.Method)
		w.Write([]byte("Hello world!"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
