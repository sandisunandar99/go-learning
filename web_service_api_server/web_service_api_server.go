package main

import (
	"encoding/json"
	"net/http"
)

type student struct {
	Name  string
	Age   int
	Class string
}

var data = []student{
	{"John", 20, "A"},
	{"Jane", 21, "B"},
	{"Bob", 22, "C"},
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// Create a new user
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)

}

func main() {
	http.HandleFunc("/users", User)
	http.ListenAndServe(":8082", nil)
}
