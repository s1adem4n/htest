package main

import (
	"flag"
	"log/slog"
	"net/http"
)

var Address = flag.String("address", "localhost:8080", "server address")

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Request", "url", r.URL.String())

		name := r.URL.Query().Get("name")

		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello, " + name + "!\n"))
	})

	slog.Info("Server started", "address", *Address)

	err := http.ListenAndServe(*Address, mux)
	if err != nil {
		panic(err)
	}
}
