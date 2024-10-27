package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./public/")))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
