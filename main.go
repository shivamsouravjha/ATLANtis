package main

import (
	"Atlantis/server"
	"Atlantis/utils"
)

func main() {
	go utils.Listener("Froms")
	go utils.Listener("Questions")
	go utils.Listener("Answers")
	go utils.Listener("Froms")
	server.Init()
}
