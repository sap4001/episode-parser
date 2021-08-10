package main

import (
	"fmt"
	"log"

	"github.com/sap4001/episode-parser/internal/server"
)

const LISTEN_PORT = 80

func main() {

	server := server.NewServer(LISTEN_PORT)
	log.Println("Starting HTTP Server on port: ", LISTEN_PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("main: unexpected error in ListenAndServe: %v", err)
	}
	fmt.Println("WWW")
}
