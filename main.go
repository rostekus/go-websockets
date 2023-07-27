package main

import (
	"log"
	"net/http"

	"github.com/rostekus/go-websockets/chat"
	"golang.org/x/net/websocket"
)

func init() {
	http.Handle("/ws", websocket.Handler(chat.HandleNewConnection))
	http.Handle("/", http.FileServer(http.Dir("./clientApp/public")))
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", nil))
}
