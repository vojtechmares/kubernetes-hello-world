package main

import (
	"log"
	"net/http"
)

var version string = "dev"

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello Kubernetes"))
}

func main() {
	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
