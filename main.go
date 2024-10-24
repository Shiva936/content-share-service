package main

import (
	"content-share/config"
	"content-share/daos"
	"content-share/router"
)

func main() {
	// Set Application Configurations
	config.Set(&config.Config{
		AppName:  "content-share-svc",
		LogLevel: "DEBUG",
		URL:      "localhost",
		Port:     "3000",
	})

	// Initialise DB (replicating DB connection, initializing variable to store in-memmory data)
	daos.InitDB()

	//Running the API Router
	router.GetRouter().Run(config.Get().URL + ":" + config.Get().Port)
}
