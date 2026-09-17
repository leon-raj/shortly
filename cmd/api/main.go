package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"shortly/internal/handler"
	"shortly/internal/store"
)

func main() {
	s, err := store.New("redirects.db")
	if err != nil {
		slog.Error("FATAL : failed to create database")
		os.Exit(1)
	}

	h := handler.New(s)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{shortCode}", h.Redirect)
	mux.HandleFunc("PUT /{shortCode}", h.ShortenURL)

	fmt.Println("Starting Server")
	err = http.ListenAndServe("0.0.0.0:2048", mux)
	println(err.Error())
}
