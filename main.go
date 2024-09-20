package main

import (
	"akro/server"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	port := flag.String("port", ":8080", "Port on which to run the server")

	go func() {
		log.Println("Start pprof on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	server.Run(*port)
}
