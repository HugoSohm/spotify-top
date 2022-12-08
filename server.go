package main

import "github.com/HugoSohm/spotify-top/src/server"

func main() {
	server.LoadEnv()
	server.StartServer()
}
