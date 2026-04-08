package main

import (
	"fmt"
)

type Ball struct {
	hits int
}

func player(name string, recv <-chan Ball, send chan<- Ball, done chan<- bool) {
	// TODO: Receive ball, increment hits, print, and send back
	// Stop when hits reaches 5
}

func main() {
	ping := make(chan Ball)
	pong := make(chan Ball)
	done := make(chan bool)

	go player("Ping", ping, pong, done)
	go player("Pong", pong, ping, done)

	// Start the game
	ping <- Ball{hits: 0}

	// Wait for game to end
	<-done
	fmt.Println("Game over!")
}
