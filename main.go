package main

import (
	"log"
	"net/http"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))

	// register new Handler
	mux.HandleFunc("/healthz", healthzHandler)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
