package chat

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type Client struct {
	id       string
	ip       string
	username string
	ws       *websocket.Conn
}

func (c *Client) StartListening() {
	buffer := make([]byte, 512)

	for {
		_, err := c.ws.Read(buffer[0:])
		if err != nil {
			ReleaseConn(c)
			exitMessage := Message{
				SenderID: "System",
				Username: "System",
				Message:  fmt.Sprintf("%s has left the chati :", c.username),
			}

			exitMessage.Broadcast()

		} else {

		}
	}
}

var clients []*Client

func HandleNewConnection(c *websocket.Conn) {
	log.Println("new connection!")
	newClient := Client{
		id:       uuid.New().String(),
		ip:       c.Request().RemoteAddr,
		username: "",
		ws:       c,
	}
	clients = append(clients, &newClient)
}

func ReleaseConn(client *Client) {
	log.Println("released conn :", client.username, client.id)
	index := -1
	for idx, val := range clients {
		if client.id == val.id {
			index = idx
		}
	}
	if index >= 0 {
		clients = append(clients[:index], clients[index+1:]...)

	}
	client.ws.Close()
}
