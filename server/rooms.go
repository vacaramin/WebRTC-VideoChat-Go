package server

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Participant Describes a single entity in hash map
type Participant struct {
	Host bool
	Conn *websocket.Conn
}

// RoomMap keeps record of each Participant in each room
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
		temp[i] = letters[rand.Intn(len(letters))]
	}
	roomID := string(temp)
	r.Map[roomID] = []Participant{}
	return roomID
}

// IntsertIntoRoom inserts a participant into a certain room
func (r *RoomMap) IntsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{host, conn}
	log.Println("Partipant added to room: ", roomID)
	r.Map[roomID] = append(r.Map[roomID], p)
}

func (r *RoomMap) DeleteRoom(roomID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	delete(r.Map, roomID)
}
