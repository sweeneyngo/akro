package main

import (
	"akro/server"
	"flag"
)

func main() {
	port := flag.String("port", ":8080", "Port on which to run the server")
	server.Run(*port)
}
