package server

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

// Init initalizes a RoomMap struct
func (r *RoomMap) Init() {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	r.Map = make(map[string][]Participant)
}

// Get Gets all the Participants in roomID
func (r *RoomMap) Get(roomId string) []Participant {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()
	return r.Map[roomId]
}

// CreateRoom Creates a New Room and returns the Room Id
func (r *RoomMap) CreateRoom() string {
	// generate a Unique ID
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

}

func (r *RoomMap) DeleteRoom() {
}
