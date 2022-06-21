package main

import (
	"unhaus/model"
	"unhaus/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}