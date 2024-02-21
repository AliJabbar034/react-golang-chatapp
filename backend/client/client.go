package client

import (
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	// Create a new WebSocket connection.
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:9000",
		Path:   "/ws",
	}
	dialer := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send a message to the server.
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, server!"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read a message from the server.
	_, message, err := conn.ReadMessage()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the message from the server.
	fmt.Println("Received:", message)
}
