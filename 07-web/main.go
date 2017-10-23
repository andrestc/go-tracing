package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.Handle("/hello", http.HandlerFunc(helloHandler))

	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
