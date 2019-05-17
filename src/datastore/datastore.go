package datastore

import (
	"context"
	"errors"
	"os"

	"google.golang.org/appengine"

	"google.golang.org/appengine/datastore"
)

// DatabaseInformation is a representation of all the information concerning
// a database, its name, the login and password
type DatabaseInformation struct {
	APIUsername  string `datastore:"API_USER_NAME"`
	APIPass      string `datastore:"API_USER_PASS"`
	InstanceName string `datastore:"INSTANCE_NAME"`
}

// findInformationFromEnv retrieves the database information from the
// environment, if one or more environment variable is missing an error is returned
func findInformationFromEnv() (DatabaseInformation, error) {
	dbInfo := DatabaseInformation{
		os.Getenv("USERNAME_DATABASE"),
		os.Getenv("PASSWORD_DATABASE"),
		os.Getenv("INSTANCE_DATABASE_NAME"),
	}
	if dbInfo.APIPass == "" || dbInfo.APIUsername == "" || dbInfo.InstanceName == "" {
		return dbInfo, errors.New("database environment variable are missing")
	}
	return dbInfo, nil
}

// findInformationFromDatastore retrieves the database information (the
// password, username and the instance name)
func findInformationFromDatastore(ctx context.Context) (DatabaseInformation, error) {
	var dbInfo DatabaseInformation
	q := datastore.NewQuery("DATABASE_INFORMATION").Limit(1)
	iterator := q.Run(ctx)

	_, err := iterator.Next(&dbInfo)

	if err != nil {
		return DatabaseInformation{}, err
	}
	return dbInfo, nil
}

// FindDatabaseInformation retrieves information from either the local
// environment or the Google Cloud Datastore, depending if we are running the
// service in dev or production
func FindDatabaseInformation(ctx context.Context) (DatabaseInformation, error) {
	var dbInfo DatabaseInformation
	var err error

	if appengine.IsDevAppServer() {
		dbInfo, err = findInformationFromEnv()
	} else {
		dbInfo, err = findInformationFromDatastore(ctx)
	}
	return dbInfo, err
}
