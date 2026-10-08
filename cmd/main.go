package main

import (
	"log"

	"github.com/M0rfes/a-go-game/game"
)

func main() {
	error := game.Run()
	if error != nil {
		log.Fatalf("failed to run %v", error)
	}
}
