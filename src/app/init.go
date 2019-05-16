package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/appengine"
)

func main() {

	if appengine.IsDevAppServer() {
		err := godotenv.Load("../../.env")
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	http.HandleFunc("/", apiHomeH)

	appengine.Main()
}

func apiHomeH(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `%s`, time.Now().Format(time.RFC850))
}
