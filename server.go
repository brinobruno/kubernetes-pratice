package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	log.Println("Starting server on :80")
	err := http.ListenAndServe(":80", nil)
	log.Println((err))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello k8s</h1>"))
}
