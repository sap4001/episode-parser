package main

import (
	"log"
	"os"
	"strconv"

	"github.com/sap4001/episode-parser/internal/server"
)

func main() {
	port := getListenPort()
	server := server.NewServer(port)
	log.Println("Starting HTTP Server on port: ", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("main: unexpected error in ListenAndServe: %v", err)
	}
}

func getListenPort() int {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("main: Listen PORT not defined: ")
	}
	i, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("main: error parsing port: %v", err)
	}
	return i
}
