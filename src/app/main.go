package main

import (
	"log"
	"net/http"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"github.com/pedafy/pedafy-tasks/src/version"
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

	if v := os.Getenv("API_VERSION"); v == "" {
		srv.SetCurrentVersion(version.Default())
	} else {
		srv.SetCurrentVersion(v)
	}

	err := srv.InitAPI()
	if err != nil {
		log.Fatal(err)
	}

	srv.RegisterHandlers()

	if len(os.Args) == 2 && os.Args[1] == "--debug" {
		err := http.ListenAndServe(":10002", nil)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}
	appengine.Main()
}
