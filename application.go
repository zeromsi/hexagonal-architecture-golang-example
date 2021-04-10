package main

import (
	"hexagonal-architecture-example/api"
	"hexagonal-architecture-example/config"
	"hexagonal-architecture-example/repository/mongo"
	"log"
)


func main() {
	config.New()
	if config.Database==string(config.MONGO){
		mongo.GetDmManager()
		log.Println(config.DatabaseConnectionString)
	}
	api.Start()
}
