package main

import "github.com/HugoSohm/spotifytop-api/src/server"

func main() {
	server.LoadEnv()
	server.StartServer()
}
