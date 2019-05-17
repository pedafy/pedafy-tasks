package apiv1

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedafy/pedafy-tasks/src/database"
	"github.com/pedafy/pedafy-tasks/src/database/layer"
	"github.com/pedafy/pedafy-tasks/src/datastore"
	"github.com/pedafy/pedafy-tasks/src/version"
)

// APIv1 represents the first version of the API
type APIv1 struct {
	Version         version.Version
	databaseHandler database.DatabaseHandler
	apiToken        string
}

// InitialisationDatabase will create a new database depending on the
// current API version
func (a *APIv1) InitialisationDatabase() {
	a.databaseHandler = layer.NewDatabaseHandler(a.Version)
}

// Register all the routes of the API to the given mux.Router
func (a *APIv1) Register(r *mux.Router) {

	a.registerMiddleware(r)

	pRouter := r.PathPrefix("/tasks/v1/").Subrouter()

	a.registerAllRoutes(r)
	a.registerAllRoutes(pRouter)
}

func (a *APIv1) registerMiddleware(r *mux.Router) {
	r.Use(a.setJSON)
}

func (a *APIv1) registerAllRoutes(r *mux.Router) {
	// Home
	r.Methods(http.MethodGet).Path("/").HandlerFunc(a.homeHandler)

	// Google App Engine
	r.Methods(http.MethodGet).Path("/_ah/start").HandlerFunc(a.startupHandler)

	// API
	// TODO: implement API here
}

func (a *APIv1) connectDatabase(ctx context.Context) {
	info, err := datastore.FindDatabaseInformation(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = a.databaseHandler.Connect(info.APIUsername, info.APIPass, "pedafy_tasks", info.InstanceName)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (a *APIv1) retrieveToken(ctx context.Context) {
	var err error
	a.apiToken, err = datastore.FindAPITokenInformation(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
