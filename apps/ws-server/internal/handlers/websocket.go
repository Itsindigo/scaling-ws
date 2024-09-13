package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader to upgrade HTTP connections to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (in production, you might want to validate the origin)
	},
}

type BaseCommand struct {
	Type string `json:"type"`
}

type JoinCommand struct {
	BaseCommand
	RoomID string `json:"roomId"`
}

func handleCommand(ws *websocket.Conn, rawMessage []byte) error {
	// First, unmarshal to the base struct to figure out the command type
	var base BaseCommand
	err := json.Unmarshal(rawMessage, &base)
	if err != nil {
		return err
	}

	// Switch based on the `type` field (discriminator)
	switch base.Type {

	case "join":
		var join JoinCommand
		err = json.Unmarshal(rawMessage, &join)
		if err != nil {
			return err
		}
		handleJoin(ws, join)

	default:
		fmt.Println("Unknown command type:", base.Type)
	}
	return nil
}

func handleJoin(ws *websocket.Conn, cmd JoinCommand) {
	fmt.Printf("User joined room: %s\n", cmd.RoomID)
	ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Joined room: %s", cmd.RoomID)))
}

func HandleWebsocketConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		// Pass the message to the command handler
		err = handleCommand(ws, msg)
		if err != nil {
			log.Println("command handling error:", err)
		}
	}
}
