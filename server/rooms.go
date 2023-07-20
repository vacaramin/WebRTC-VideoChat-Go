package server

import (
	"math/rand"
	"sync"
	"time"

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
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	temp := make([]rune, 8)
	for i := range temp {
		temp[i] := letters[rand.Intn(len(letters))]
	}
	roomID := string(b)
	r.Map[roomID] = []Participant{}
	return roomID
}

func (r *RoomMap) DeleteRoom() {
}
