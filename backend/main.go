package main

import (
	"log"
	"os"

	"github.com/WibuSOS/sinarmas/api"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.SetOutput(file)

	db, err := api.SetupDb()
	if err != nil {
		log.Panicln(err.Error())
	}

	server := api.MakeServer(db)
	server.RunServer()
}
