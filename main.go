package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	log.Println("starting mrcp proxy http command on :7890")

	err := http.ListenAndServe(":7890", mux)

	log.Fatal(err)
}
