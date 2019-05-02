package pedafytasks

import (
	"fmt"
	"net/http"
	"time"
)

func init() {

	http.HandleFunc("/", apiHomeH)
}

func apiHomeH(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	fmt.Fprintf(w, `{"running":true, "time":"%s"}`, time.Now().String())
}
