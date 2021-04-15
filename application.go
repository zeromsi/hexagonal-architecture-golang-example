package main

import (
	"hexagonal-architecture-example/api"
	"hexagonal-architecture-example/config"
	"hexagonal-architecture-example/repository/mongo"
)


func main() {
	config.New()
	if config.Database==string(config.MONGO){
		mongo.GetDmManager()
	}else if config.Database==string(config.MYSQL){
 		// Initialize mysql connection
	}else if config.Database==string(config.REDIS){
		// Initialize redis connection
	}
	api.Start()
}
