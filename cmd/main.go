package main


import "redisServer/internal/server"

func main() {
	server.RunIoMultiplexingServer()
}