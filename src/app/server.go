package main

import (
	"github.com/pedafy/pedafy-tasks/src/version"
)

type server struct {
	currentVersion version.Version
}

func (s *server) InitAPI() error {
	// do things here
	return nil
}

func (s *server) RegisterHandlers() {
	// do things here
}
