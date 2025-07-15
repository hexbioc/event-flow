package main

import "collector/server"

func main() {
	server, shutdownActions := server.New()
	defer shutdownActions()

	server.Run()
}
