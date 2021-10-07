package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from server"))
	})
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err.Error())
	}
}
