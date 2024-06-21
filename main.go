package main

import (
	"flag"
	"net/http"
)

var Address = flag.String("address", "localhost:8080", "server address")

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello\n"))
	})
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Path + "\n"))
	})
	mux.HandleFunc("/banana", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("banana\n"))
	})

	err := http.ListenAndServe(*Address, mux)
	if err != nil {
		panic(err)
	}
}
