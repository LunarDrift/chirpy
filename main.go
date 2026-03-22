package main

import (
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
