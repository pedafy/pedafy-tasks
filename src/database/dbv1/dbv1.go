package dbv1

import (
	"database/sql"
	"fmt"

	"google.golang.org/appengine"
)

// DBv1 refers to the first version of the database interface
type DBv1 struct {
	db *sql.DB
}

// Connect the database to a MySQL server using the given
// credentials and URL
func (d *DBv1) Connect(user, pass, dbName, url string) error {
	var err error

	if appengine.IsDevAppServer() {
		d.db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp([127.0.0.1]:3306)/%s?parseTime=true", user, pass, dbName))
	} else {
		d.db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/%s?parseTime=true", user, pass, url, dbName))
	}

	if err == nil {
		err = d.db.Ping()
	}
	return err
}

// IsNew returns true if the database is not connected yet
func (d *DBv1) IsNew() bool {
	return d.db == nil
}
