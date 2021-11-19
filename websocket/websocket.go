package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"youtube/youtube"

	"github.com/gorilla/websocket"
	// "golang.org/x/net/websocket"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// fmt.Fprintf(w, "&{}", upgrader.CheckOrigin)
	ws, err := upgrader.Upgrade(w, r, nil)
	fmt.Println(w, ws)
	if err != nil {
		fmt.Println(err)
		// fmt.Fprint(w, "hello websocket rgrfnb bngt tgh")
		return ws, err
	}
	return ws, nil
}
func Writer(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(5 * time.Second)
		for t := range ticker.C {
			fmt.Printf("updating status : %+v\n", t)
			items, err := youtube.GetSubscriber()
			if err != nil {
				fmt.Println("err")
			}
			jsonString, err := json.Marshal(items)
			if err != nil {
				fmt.Println(err)
			}
			// fmt.Println("jhghjgjhsd")
			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
				fmt.Println(err)
				return
			}

		}
	}
}
