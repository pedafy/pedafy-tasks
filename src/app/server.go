package main

import (
	"errors"
	"net/http"

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

	// TODO: register handlers using gorilla mux

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello")); err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
		}
	})
}
