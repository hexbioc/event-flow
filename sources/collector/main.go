package main

import "collector/server"

func main() {
	server := server.New()

	server.Run()
}
