package main

import (
	"log"
	"github.com/somraj2929/week2-GL1-CipherSchools/database"
	"github.com/somraj2929/week2-GL1-CipherSchools/routers"
)

func main() {
	database.Setup()
	engine := routers.Router()
	err := engine.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}