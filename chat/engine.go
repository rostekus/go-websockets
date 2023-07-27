package chat

import (
	"encoding/json"
	"fmt"
)

func HandleInputMessage(from *Client, data []byte) {

	var input map[string]string
	json.Unmarshal(data, &input)

	switch input["action"] {
	case "post_message":
		newMessage := Message{
			SenderID: from.id,
			Username: from.username,
			Message:  input["message"],
		}
		newMessage.Post()
	case "initial_message":
		from.username = input["username"]
		newMessage := Message{
			SenderID: "System",
			Username: "System",
			Message:  fmt.Sprintf("%x joined the chat", from.username),
		}
		newMessage.BroadcastTo(from)

		chatHistory, _ := json.Marshal(messages)
		from.ws.Write(chatHistory)
	}
}
