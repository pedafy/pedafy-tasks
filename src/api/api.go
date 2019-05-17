package api

import "github.com/gorilla/mux"

// APIHandler interfaces all API versions
type APIHandler interface {
	InitialisationDatabase()
	Register(r *mux.Router)
}
