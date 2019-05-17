package apiv1

import (
	"log"
	"net/http"
	"time"

	"google.golang.org/appengine"
)

func (a *APIv1) startupHandler(w http.ResponseWriter, r *http.Request) {
	if a.databaseHandler.IsNew() {
		a.connectDatabase(appengine.NewContext(r))
	}
	if a.apiToken == "" {
		a.retrieveToken(appengine.NewContext(r))
	}
	if _, err := w.Write([]byte("ready to operate")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err.Error())
	}
}

func (a *APIv1) homeHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(time.Now().Format(time.RFC850))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err.Error())
	}
}

func (a *APIv1) setJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		next.ServeHTTP(w, r)
	})
}
