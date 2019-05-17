package dbv1

import (
	"database/sql"
	"strconv"

	"github.com/pedafy/pedafy-tasks/src/database"
)

// GetAllStatus will return an array of status from the database
func (d *DBv1) GetAllStatus() ([]database.Status, error) {
	var s = "SELECT * FROM `status`"
	resp, err := d.db.Query(s)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	status := make([]database.Status, 0)

	for resp.Next() {
		var curr database.Status
		err = resp.Scan(&curr.ID, &curr.Name)
		if err != nil {
			return nil, err
		}
		status = append(status, curr)
	}
	return status, nil
}

// GetStatusByID return the status matching the given ID
func (d *DBv1) GetStatusByID(ID int) (database.Status, error) {
	var s = "SELECT * FROM `status` WHERE id = ?"
	resp, err := d.db.Query(s, strconv.Itoa(ID))
	if err != nil {
		return database.Status{}, err
	}
	defer resp.Close()

	return d.oneRowQuery(resp)
}

// GetStatusByName get one single row from the database matching the given
// name
func (d *DBv1) GetStatusByName(name string) (database.Status, error) {
	var s = "SELECT * FROM `status` WHERE name = ?"
	resp, err := d.db.Query(s, name)
	if err != nil {
		return database.Status{}, err
	}
	defer resp.Close()

	return d.oneRowQuery(resp)
}

func (d *DBv1) oneRowQuery(resp *sql.Rows) (database.Status, error) {
	var err error
	var status database.Status

	if resp.Next() {
		err = resp.Scan(&status.ID, &status.Name)
		if err != nil {
			return database.Status{}, err
		}
	} else {
		return status, nil
	}
	return status, nil
}
