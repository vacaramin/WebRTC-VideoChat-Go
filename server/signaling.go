package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var AllRooms RoomMap

// CreateRoomRequestHandler Create a Room and return Room ID
func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	roomID := AllRooms.CreateRoom()
	type resp struct {
		RoomID string `json:"room_id"`
	}
	log.Println(AllRooms.Map)
	json.NewEncoder(w).Encode(resp{RoomID: roomID})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type broadcastMsg struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}

var broadcast = make(chan broadcastMsg)

func broadcaster() {
	for {
		msg := <-broadcast

		for _, client := range AllRooms.Map[msg.RoomID] {
			if client.Conn != msg.Client {
				err := msg.Client.WriteJSON(msg.Message)
				if err != nil {
					log.Fatal("JSON ERROR", err)
					client.Conn.Close()
				}
			}
		}
	}
}

// JoinRoomRequest Handler Will join the client in a particular room
func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {

	roomID, ok := r.URL.Query()["roomID"]
	if !ok {
		log.Printf("roomID MISSING in URL")
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Websocket Error", err)
	}
	AllRooms.IntsertIntoRoom(roomID[0], false, ws)

	go broadcaster()

	for {
		var msg broadcastMsg
		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatal("Read Error", err)
		}
		msg.Client = ws
		msg.RoomID = roomID[0]
		broadcast <- msg
	}
}
