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
		w.Write([]byte("hello"))
	})
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Path))
	})

	err := http.ListenAndServe(*Address, mux)
	if err != nil {
		panic(err)
	}
}
