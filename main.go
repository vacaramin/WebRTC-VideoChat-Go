package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/create", CreateRoomRequestHandler)
	http.Handle("/join", JoinRoomRequestHandler)

	log.Println("Server started on Port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
