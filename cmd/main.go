package main

import (
	"Portfolio_You/server"
	"log"
)

func main() {
	// if err := config.Init(); err != nil {
	// 	log.Fatalf("%s", err.Error())
	// }

	app := server.NewApp()

	if err := app.Run("8000"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
