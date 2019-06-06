package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pedafy/pedafy-tasks/src/api"
	"github.com/pedafy/pedafy-tasks/src/api/layer"
	"github.com/pedafy/pedafy-tasks/src/version"
)

type server struct {
	currentVersion version.Version
	apiHandler     api.APIHandler
}

// SetCurrentVersion will set the version to the given value
func (s *server) SetCurrentVersion(currentVersion version.Version) {
	s.currentVersion = currentVersion
}

// InitAPI will create the API depending on the server's main version
func (s *server) InitAPI() error {
	if s.currentVersion == "" {
		return errors.New("no API version")
	}
	s.apiHandler = layer.NewAPIHandler(s.currentVersion)
	s.apiHandler.InitialisationDatabase()
	return nil
}

// RegisterHandlers will register all API URL
func (s *server) RegisterHandlers() {

	r := mux.NewRouter()

	s.apiHandler.Register(r)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}
