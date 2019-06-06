package datastore

import (
	"context"
	"errors"
	"os"

	"google.golang.org/appengine"

	"google.golang.org/appengine/datastore"
)

// TokenAPI is the data structure fitting google cloud datastore
type TokenAPI struct {
	Token string `datastore:"TOKEN_VALUE"`
}

// findTokenFromEnv retrieves the API token from the environment,
// if one or more environment variable is missing an error is returned
func findTokenFromEnv() (string, error) {
	token := os.Getenv("TOKEN_VALUE")
	if token == "" {
		return "", errors.New("api token variable is missing")
	}
	return token, nil
}

// findTokenFromDatastore retrieves the API token
func findTokenFromDatastore(ctx context.Context) (string, error) {
	var info TokenAPI
	q := datastore.NewQuery("API_TASKS_TOKEN").Limit(1)
	iterator := q.Run(ctx)

	_, err := iterator.Next(&info)

	if err != nil {
		return "", err
	}
	return info.Token, nil
}

// FindAPITokenInformation retrieves information about the API token
// from either the local environment or the Google Cloud Datastore,
// depending if we are running the service in dev or production
func FindAPITokenInformation(ctx context.Context) (string, error) {
	var token string
	var err error

	if appengine.IsDevAppServer() {
		token, err = findTokenFromEnv()
	} else {
		token, err = findTokenFromDatastore(ctx)
	}
	return token, err
}
