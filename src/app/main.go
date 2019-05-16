package main

import (
	"log"

	"github.com/joho/godotenv"
	"google.golang.org/appengine"
)

func main() {
	var srv server

	if appengine.IsDevAppServer() {
		err := godotenv.Load("../../.env")
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	err := srv.InitAPI()
	if err != nil {
		log.Fatal(err.Error())
	}

	srv.RegisterHandlers()

	appengine.Main()
}
