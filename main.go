package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func reader(conn *websocket.Conn) {
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(msg))
		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage HTTP!")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the WebSocket Endpoint!")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.Println("Client Connected")
	err = ws.WriteMessage(websocket.TextMessage, []byte("Hello Client"))
	if err != nil {
		log.Println(err)
	}
	reader(ws)
}

func setupRoutes() {
	fmt.Println("Setting up routes...")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}
func main() {
	fmt.Println("Hello WebSocket Server!")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
