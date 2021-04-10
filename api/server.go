package api

import (
	"hexagonal-architecture-example/config"
	"log"
	"net/http"
)
func Start() {
	Routes()
	log.Println("Server starting on port: ",config.ServerPort)
	err := http.ListenAndServe(":"+config.ServerPort, nil)
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server on port %s\n", config.ServerPort)
	}
}
