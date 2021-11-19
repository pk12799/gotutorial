package main

import (
	"fmt"
	"log"
	"net/http"

	// "youtube/websocket"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsENdpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// log.Printf("hfgh fgfhgsd gfdgsdg hgd")
	Reader(ws)
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloi")
}

func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
		}
	}
}

// func wsENdpoint(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "websockwt")
// }

func setupRoutes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/ws", wsENdpoint)
}

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
