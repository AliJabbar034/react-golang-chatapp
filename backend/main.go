// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// func main() {
// 	mux := mux.NewRouter()
// 	mux.HandleFunc("/ws", Listen)

// 	http.ListenAndServe(":9000", mux)
// }

// func Listen(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		panic(err)
// 		return
// 	}
// 	for {
// 		messageType, message, err := conn.ReadMessage()
// 		if err != nil {
// 			panic(err)
// 			return
// 		}
// 		fmt.Println("messageType", message)
// 		er := conn.WriteMessage(messageType, []byte("message"))
// 		if er != nil {
// 			panic(er)
// 			return
// 		}
// 	}

// }

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	// Create a new WebSocket server.
	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// Allow all connections
			return true
		},
	}
	type Res struct {
		Type string `json:"type"`
		Body string `json:"body"`
	}
	// Listen for WebSocket connections on port 8080.
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade the HTTP connection to a WebSocket connection.

		conn, err := wsUpgrader.Upgrade(w, r, nil)
		defer conn.Close()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Read messages from the client.
		for {
			// Read a message from the client.
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			// Print the message to the console.
			fmt.Println("Received:", message)
			// Send a message back to the client.
			err = conn.WriteMessage(messageType, []byte("message"))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	})
	// Start the server.
	http.ListenAndServe(":9000", nil)
}
