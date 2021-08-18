package main

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/test", test)

	http.ListenAndServe(":8088", nil)
}
