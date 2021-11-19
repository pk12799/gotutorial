package main

import (
	"fmt"
	"log"
	"net/http"
	"youtube/websocket"
)

func stats(w http.ResponseWriter, r *http.Request) {

	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
		fmt.Fprint(w, "hello web")
	}
	fmt.Println(ws)
	go websocket.Writer(ws)
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
func setupRoute() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func main() {
	// item, err := youtube.GetSubscriber()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // websocket.Writer()
	// fmt.Printf("%+v\n", item)

	setupRoute()
}
