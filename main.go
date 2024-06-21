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

	http.ListenAndServe(*Address, mux)
}
