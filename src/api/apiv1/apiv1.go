package apiv1

import (
	"github.com/gorilla/mux"
	"github.com/pedafy/pedafy-tasks/src/database"
	"github.com/pedafy/pedafy-tasks/src/database/layer"
	"github.com/pedafy/pedafy-tasks/src/version"
)

// APIv1 represents the first version of the API
type APIv1 struct {
	version         version.Version
	databaseHandler database.DatabaseHandler
}

// InitialisationDatabase will create a new database depending on the
// current API version
func (a *APIv1) InitialisationDatabase() {
	a.databaseHandler = layer.NewDatabaseHandler(a.version)
}

// Register all the routes of the API to the given mux.Router
func (a *APIv1) Register(r *mux.Router) {

}
