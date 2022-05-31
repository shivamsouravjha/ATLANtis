package main

import (
	"Atlantis/server"
	"Atlantis/utils"
)

func main() {
	go utils.Listener("Forms")
	go utils.Listener("Questions")
	go utils.Listener("Answers")
	go utils.Listener("Responses")
	server.Init()
}
