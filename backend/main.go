package main

import (
	"log"
	"os"

	"github.com/WibuSOS/sinarmas/backend/api"
	"github.com/WibuSOS/sinarmas/backend/database"
	"github.com/WibuSOS/sinarmas/backend/utils/localization"

	_ "github.com/joho/godotenv/autoload"
)

func init() {
	go localization.WriteJSON()
}

func main() {
	file, err := os.OpenFile("./logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.SetOutput(file)

	db, err := database.SetupDb()
	if err != nil {
		log.Panicln(err.Error())
	}

	server := api.MakeServer(db)
	server.RunServer()
}
