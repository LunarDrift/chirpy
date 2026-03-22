package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
